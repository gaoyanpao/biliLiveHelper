package biliLiveHelper

type CmdType string

const (
	CmdAll                   CmdType = ""
	CmdLive                  CmdType = "LIVE"
	CmdPreparing             CmdType = "PREPARING"
	CmdDanmuMsg              CmdType = "DANMU_MSG"
	CmdWelcomeGuard          CmdType = "WELCOME_GUARD"
	CmdWelcome               CmdType = "WELCOME"
	CmdSendGift              CmdType = "SEND_GIFT"
	CmdNoticeMsg             CmdType = "NOTICE_MSG"
	CmdRealTimeMessageUpdate CmdType = "ROOM_REAL_TIME_MESSAGE_UPDATE"
	CmdOnlineChange          CmdType = "ONLINE_CHANGE"
)

type MsgHandler interface {
	Handle(ctx *Context) bool
}

type Handle func(ctx *Context) bool

func (f Handle) Handle(ctx *Context) bool { return f(ctx)}

type HandleChain []MsgHandler

func (c *Client) RegHandleFunc(cmdType CmdType, handler Handle) {
	if len(c.handlers) >= abortIndex {
		panic("too many handlers")
	}
	c.handlers[cmdType] = append(c.handlers[cmdType], handler)
}

