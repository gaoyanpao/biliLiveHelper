package biliLiveHelper

const (

	BiliUrl = "https://api.live.bilibili.com"

	//获取一个房间的基本信息
	URLMobileRoomInit = "/room/v1/Room/mobileRoomInit"

	//获取一个房间的详细信息
	URLGetInfo = "/room/v1/Room/get_info"

	//进入房间时客户端将访问该接口
	//访问该接口将在自己的账户中产生一条观看直播的历史记录
	URLRoomEntryAction = "/room/v1/Room/room_entry_action"

	//获取弹幕服务器
	URLGetDanmuConf = "/room/v1/Danmu/getConf"

	//获取主播的头像和等级一类的信息
	URLGetAnchorInRoom = "/live_user/v1/UserInfo/get_anchor_in_room"

	//获取自己在直播站的基本信息
	URLGetUser = "/mobile/getUser"

	//获取自己在当前直播间的信息
	URLGetInfoInRoom = "/live_user/v1/UserInfo/get_info_in_room"

	//获取所有头衔
	URLGetTitle = "/appUser/getTitle"

	//查询是否关注了当前主播
	URLIsFollowed = "/relation/v1/Feed/isFollowed"

	//获取上方的 Tab. 互动, 主播, 贡献榜
	URLMobileTab = "/room/v2/Room/mobileTab"

	//获取房间的历史弹幕(10条)
	URLRoomMessage = "/AppRoom/msg"

	//获取进房后右下角banner
	URLMobileRoomBanner = "/activity/v1/Common/mobileRoomBanner"

	//获取各种礼物的基本信息
	URLGetGiftConfig = "/gift/v3/live/gift_config"

	//获取访问小时总榜
	URLRoomRank = "/rankdb/v1/Common/roomRank"

	//直播站首页
	URLGetIndexList = "/xlive/app-interface/v2/index/getAllList"

	//获取某个分类下的全部子分类
	URLGetAreaList = "/room/v1/Area/getList"

	//根据某种维度来获取房间列表
	URLGetRoomList = "/room/v3/Area/getRoomList"

	//发送弹幕
	URLSendMessage = "/api/sendmsg"
)
