package biliLiveHelper

import (
	"errors"
	"github.com/gaoyanpao/biliLiveHelper/tools"
	"strconv"
)

/**
 * 获取一个房间的基本信息
 * /room/v1/Room/mobileRoomInit
 * @param id 房间号或房间短号
 */
func MobileRoomInit(roomID int64) {}

/**
 * 获取一个房间的详细信息
 * /room/v1/Room/get_info
 * @param id 房间号
 */
func GetInfo(roomID int) (*RoomInfo, error) {
	resp := &GetInfoResp{}
	err := tools.Get(BiliUrl+URLGetInfo, map[string]string{"id": strconv.Itoa(roomID),}, nil, resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, err
}

/**
 * 进入房间时客户端将访问该接口
 * 访问该接口将在自己的账户中产生一条观看直播的历史记录
 * /room/v1/Room/room_entry_action
 * @param roomId 房间号(没试过能不能用短号, 下同)
 */
func RoomEntryAction(roomID int64) {}

/**
 * 获取弹幕服务器
 * /room/v1/Danmu/getConf
 * @param roomId 房间号
 */
func GetDanmuConf(roomID int) (*GetConfData, error) {
	resp := &GetConfResp{}
	err := tools.Get(BiliUrl+URLGetDanmuConf, map[string]string{"room_id": strconv.Itoa(roomID),}, nil, resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, err
}

/**
 * 获取该房间的主播的头像和等级一类的信息
 * /live_user/v1/UserInfo/get_anchor_in_room
 * @param roomId 房间号
 */
func GetAnchorInRoom(roomID int64) {}

/**
 * 获取自己在直播站的基本信息, 包括自己的直播间号, 银瓜子, 金瓜子数量等
 * /mobile/getUser
 */
func GetUser() {}

/**
 * 获取自己在当前直播间的信息, 包括自己的权限以及是否是管理员等
 * /live_user/v1/UserInfo/get_info_in_room
 * @param roomId 房间号
 */
func GetInfoInRoom(roomID int64) {}

/**
 * 获取所有头衔
 * /appUser/getTitle
 * @param scale 屏幕尺寸
 */
func GetTitle() {}

/**
 * 查询是否关注了当前主播
 * /relation/v1/Feed/isFollowed
 * @param follow 所查询的主播的用户 ID
 */
func IsFollowed(roomID int64) {}

/**
 * 进入直播间的时候, 客户端会访问该接口来动态获取上方的 Tab. 包括 互动, 主播, 贡献榜 等
 * /room/v2/Room/mobileTab
 * @param roomId 房间号
 */
func MobileTab(roomID int64) {}

/**
 * 获取房间的历史弹幕(10条)
 * /AppRoom/msg
 * @param roomId 房间号
 */
func RoomMessage(roomID int64) {}

/**
 * 获取进房后右下角显示的那些东西, 通常是一些活动, 它们导向 H5 页面
 * /activity/v1/Common/mobileRoomBanner
 * @param roomId 房间号
 * @param roomUserId 主播的用户 ID
 */
func MobileRoomBanner(roomID int64) {}

/**
 * 获取各种礼物的基本信息, 包括贴图地址, 描述, 价格等
 * /gift/v3/live/gift_config
 */
func GetGiftConfig(roomID int64) {}

/**
 * 获取访问 小时总榜 的地址(H5)
 * /rankdb/v1/Common/roomRank
 */
func RoomRank(roomID int64) {}

/**
 * 直播站首页
 * 首页 -> 直播
 * /xlive/app-interface/v2/index/getAllList
 */
func GetIndexList() {}

/**
 * 获取某个直播分类下的全部子分类
 * /room/v1/Area/getList
 */
func GetAreaList() {}

/**
 * 根据某种维度来获取房间列表
 * area, parent, category 为 0 表示不筛选这些维度
 * sortType 为 null 表示不排序
 *
 * 首页 -> 直播 -> 查看更多/全部直播
 * /room/v3/Area/getRoomList
 * @param page 分页, 从 1 开始
 * @param sortType 排序维度, 已知的有 online(最热直播), live_time(最新开播)
 */
func GetRoomList() {}

/**
 * 发送弹幕(直播)
 * /api/sendmsg
 * @param bubble 气泡, 不明确含义
 * @param cid 房间号
 * @param mid 发送者的用户 ID
 * @param message 弹幕内容
 * @param random 随机数, 不包括符号位有 9 位 或者 10 位
 * @param mode 弹幕模式, 可能与视频弹幕的模式含义相同, 可能需要特殊身份才能使用额外模式, 下同
 * @param pool 弹幕池
 * @param type 固定为 "json"
 * @param color 弹幕颜色
 * @param fontSize 弹幕字号
 * @param playTime 不明确
 */
func SendMessage() {}
