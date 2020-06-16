package biliLiveHelper

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gaoyanpao/biliLiveHelper/tools"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	SimpleRoomInfo
	connected bool
	Conn      *websocket.Conn
	msgChan   chan []byte
	handlers  map[CmdType]HandleChain
}

func NewClient(roomID int) *Client {
	roomInfo, err := GetInfo(roomID)
	if err != nil {
		print(err.Error())
		return nil
	}
	client := &Client{}
	client.SimpleRoomInfo = roomInfo.SimpleRoomInfo
	client.msgChan = make(chan []byte)
	client.handlers = make(map[CmdType]HandleChain)
	return client
}

func (c *Client) PrintRoomInfo() {
	log.Printf("房间: %s[%d]\n关注: %d\n人气: %d\n直播状态: %v",
		c.SimpleRoomInfo.Title,
		c.SimpleRoomInfo.RoomID,
		c.SimpleRoomInfo.Attention,
		c.SimpleRoomInfo.Online,
		getLiveStatusString(c.SimpleRoomInfo.LiveStatus))
}

func (c *Client) StartListen() error {
	conf, err := GetDanmuConf(c.RoomID)
	if conf == nil || err != nil {
		return errors.New("获取弹幕服务器配置失败")
	}
	hostCont := len(conf.HostServerList)
	if hostCont == 0 {
		return errors.New("未找到弹幕服务器")
	}
	host := conf.HostServerList[hostCont-1]
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", host.Host, host.WsPort), Path: "/sub"}
	var wsClient *websocket.Dialer
	headers := http.Header{}
	headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	conn, _, err := wsClient.Dial(u.String(), headers)
	if err != nil {
		return err
	}
	c.Conn = conn
	// 发送进入房间数据包
	sendData := map[string]interface{}{"uid": 0, "roomid": c.RoomID, "protover": 1, "platform": "web", "clientver": "1.12.0", "type": 2, "key": conf.Token}
	packet := NewPacket(1, OpEnterRoomSend)
	err = packet.WriteJSONBody(sendData)
	if err != nil {
		return err
	}
	c.Send(packet)
	c.connected = true
	go c.heartbeatLoop()
	go c.handleMsgLoop()
	c.ReceiveMsgLoop()
	return nil
}

func (c *Client) ReceiveMsgLoop() {
	for c.connected {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v\n", err)
			}
			break
		}
		c.msgChan <- message
	}
	log.Println("与弹幕服务器连接断开")
	c.connected = false
}

func (c *Client) IsConnected() bool {
	return c.connected
}

func (c *Client) Send(packet *DataPacket) {
	if err := c.Conn.WriteMessage(websocket.BinaryMessage, packet.Bytes()); err != nil {
		print("send error: ", err.Error())
	}
}

func (c *Client) handleMsgLoop() {
	for {
		message := <-c.msgChan
		packets, err := PacketFromBytes(message)
		if err != nil {
			log.Println("数据包解析错误")
			break
		}
		for _, packet := range packets {
			switch packet.Operation {
			case OpEnterRoomRecv:
				handleEnterRoomRecv(packet)
			case OpHeartbeatRecv:
				c.Online = binary.BigEndian.Uint32(packet.Body[:4])
			case OpNotice:
				c.handleNotice(packet)
			}
		}
	}
}

func (c *Client) heartbeatLoop() {
	packet := NewHeartBeatPacket()
	for c.connected {
		c.Send(packet)
		time.Sleep(time.Second * 30)
	}
	log.Print("与弹幕服务器连接断开")
}

func handleEnterRoomRecv(packet *DataPacket) {
	data := &EnterRoomRecv{}
	jsoniter.Unmarshal(packet.Body, data)
	if data.Code == 0 {
		fmt.Println("成功连接弹幕服务器...")
	}
}

func (c *Client) handleNotice(packet *DataPacket) {
	if packet.ProtocolVersion == 0 || packet.ProtocolVersion == 1 {
		data, err := simplejson.NewJson(packet.Body)
		if err != nil {
			panic(err)
		}
		//jsoniter.Unmarshal(packet.Body, data)
		cmdType := CmdType(data.Get("cmd").MustString())
		c.RunCmdHandlers(NewContext(cmdType, data))
		//switch cmdType {
		//case CmdDanmuMsg:
		//	log.Printf("[弹幕]<%v>%s", data.Get("info").GetIndex(2).GetIndex(1).MustString(), data.Get("info").GetIndex(1).MustString())
		//case CmdWelcome:
		//	log.Printf("[欢迎]%v", data.Get("data"))
		//case CmdSendGift:
		//	log.Printf("[礼物]%v", data.Get("data"))
		//default:
		//	log.Printf("\033[1;;31m忽略的消息类型: %s\033[0m\n", cmdType)
		//}
	} else {
		dataList, err := PacketFromBytes(tools.DoZlibUnCompress(packet.Body))
		if err != nil {
			return
		}
		for _, d := range dataList {
			c.handleNotice(d)
		}
	}
}

func (c *Client) RunCmdHandlers(ctx *Context) {
	if handlers, ok := c.handlers[ctx.Cmd]; ok {
		if handlers != nil {
			ctx.handlers = handlers
			ctx.Next()
		}
	}
}

func getLiveStatusString(status int) string {
	switch status {
	case 0:
		return "未开播"
	case 1:
		return "直播中"
	case 2:
		return "轮播中"
	default:
		return fmt.Sprintf("未知状态%d", status)
	}
}
