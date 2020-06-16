package biliLiveHelper

var HeartBeatBody = []byte("[object Object]")

type DefaultResp struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
}

type BiliResp interface {
	IsSuccess() 	bool
	GetMessage() 	string
	GetCode() 		int
	GetMsg() 		string
}

type GetConfResp struct {
	DefaultResp
	Data GetConfData `json:"data"`
}

//GetConf

type GetConfData struct {
	Token            string       `json:"token"`
	RefreshRowFactor float32      `json:"refresh_row_factor"`
	RefreshRate      int          `json:"refresh_rate"`
	MaxDelay         int          `json:"max_delay"`
	Port             int          `json:"port"`
	Host             string       `json:"host"`
	HostServerList   []ServerData `json:"host_server_list"`
	ServerList       []ServerData `json:"server_list"`
}

type ServerData struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	WssPort int    `json:"wss_port"`
	WsPort  int    `json:"ws_port"`
}


//GetRoomInfo

type GetInfoResp struct {
	DefaultResp
	Data RoomInfo `json:"data,omitempty"`
}

type SimpleRoomInfo struct {
	UID string `json:"uid"`
	RoomID int `json:"room_id"`
	ShortID int `json:"short_id"`
	Attention int `json:"attention"`
	Online uint32 `json:"online"`
	Description string `json:"description"`
	LiveStatus int `json:"live_status"`
	AreaID int `json:"area_id"`
	AreaName string `json:"area_name"`
	ParentAreaID int `json:"parent_area_id"`
	ParentAreaName string `json:"parent_area_name"`
	OldAreaID int `json:"old_area_id"`
	Title string `json:"title"`
	LiveTime string `json:"live_time"`
}

type RoomInfo struct {
	SimpleRoomInfo
	IsPortrait bool `json:"is_portrait"`
	AreaPendants string `json:"area_pendants"`
	Background string `json:"background"`
	UserCover string `json:"user_cover"`
	Keyframe string `json:"keyframe"`
	IsStrictRoom bool `json:"is_strict_room"`
	Tags string `json:"tags"`
	IsAnchor int `json:"is_anchor"`
	RoomSilentType int `json:"room_silent_type"`
	RoomSilentLevel int `json:"room_silent_level"`
	RoomSilentSecond int `json:"room_silent_second"`
	HotWords []string `json:"hot_words"`
	HotWoRdsStatus int `json:"hot_wrods_status"`
	Verify string `json:"verify"`
	NewPendants Pendants `json:"new_pendants"`
	UpSession string `json:"up_session"`
	PKStatus int `json:"pk_status"`
	PKID int `json:"pk_id"`
	BattleID int `json:"battle_id"`
	AllowChangeAreaTime int `json:"allow_change_area_time"`
	AllowChangeCoverTime int `json:"allow_change_cover_time"`
	StudioInfo StudioInfo `json:"studio_info"`
}

type Pendants struct {
	Frame FrameData `json:"frame"`
	Badge interface{} `json:"badge"`
	MobileFrame FrameData `json:"mobile_frame"`
	MobileBadge interface{} `json:"mobile_badge"`
}

type FrameData struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Position int `json:"position"`
	Desc string `json:"desc"`
	Area int `json:"area"`
	AreaOld int `json:"area_old"`
	BgColor string `json:"bg_color"`
	BgPic string `json:"bg_pic"`
	UseOldArea bool `json:"use_old_area"`
}

type StudioInfo struct {
	Status int `json:"status"`
	MasterList []interface{} `json:"master_list"`
}