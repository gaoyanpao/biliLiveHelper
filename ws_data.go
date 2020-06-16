package biliLiveHelper

// 发送进房数据包后的回应
type EnterRoomRecv struct {
	Code int `json:"code"`
}


type NoticeRecv struct {
	CMD  string      `json:"cmd"`
	Info []interface{} `json:"info"`
	Data interface{} `json:"data"`
}

type NoticeData struct {
	UID        int  `json:"uid"`
	UserName   int  `json:"uname"`
	IsAdmin    bool `json:"id_admin"`
	SVip       int  `json:"svip"`
	Vip        int  `json:"vip"`
	MockEffect int  `json:"mock_effect"`
}

type DanmuMsg struct {
	UserName string
	Content  string
}
