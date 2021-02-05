package biliLiveHelper

import "github.com/bitly/go-simplejson"

type Gift struct {
	UserName string
	GiftName string
	Action   string
}

func ParseGift(jsonData *simplejson.Json) *Gift {
	data := jsonData.Get("data")
	return &Gift{
		UserName: data.Get("uname").MustString(),
		GiftName: data.Get("giftName").MustString(),
		Action:   data.Get("action").MustString(),
	}
}
