package biliLiveHelper

type Danmu struct {
	Text string `json:"text"`
	Nickname string `json:"nickname"`
	UID int `json:"uid"`
	UserNameColor string `json:"uname_color"`
	Timeline string `json:"timeline"`
	IsAdmin string `json:"is_admin"`
	Vip int `json:"vip"`
	SVip int `json:"svip"`
	Medal []interface{} `json:"medal"`
	Title []string `json:"title"`
}