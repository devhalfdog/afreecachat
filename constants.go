package afreecachat

// Reference : https://github.com/wakscord/afreeca
const (
	SVC_KEEPALIVE             = 0
	SVC_LOGIN                 = 1
	SVC_JOINCH                = 2
	SVC_QUITCH                = 3
	SVC_CHUSER                = 4
	SVC_CHATMESG              = 5
	SVC_SETCHNAME             = 6
	SVC_SETBJSTAT             = 7
	SVC_SETDUMB               = 8
	SVC_DIRECTCHAT            = 9
	SVC_NOTICE                = 10
	SVC_KICK                  = 11
	SVC_SETUSERFLAG           = 12 // "\x1b\t001200006500\f720928|163840\ftjdwns342(2)\f하나만물어봄\f0\f0\f589856|163840\f"
	SVC_SETSUBBJ              = 13
	SVC_SETNICKNAME           = 14 // "\x1b\t001400006700\fzlah0864\f종우패는곰이\f1\f269025824|163840\f종우패는호종\f"
	SVC_SVRSTAT               = 15
	SVC_RELOADHOST            = 16
	SVC_CLUBCOLOR             = 17
	SVC_SENDBALLOON           = 18
	SVC_ICEMODE               = 19
	SVC_SENDFANLETRTRER       = 20
	SVC_ICEMODE_EX            = 21
	SVC_GET_ICEMODE_RELAY     = 22
	SVC_SLOWMODE              = 23
	SVC_RELOADBURNLEVEL       = 24
	SVC_BLINDKICK             = 25
	SVC_MANAGERCHAT           = 26
	SVC_APPENDDATA            = 27
	SVC_BASEBALLEVENT         = 28
	SVC_PAIDITEM              = 29
	SVC_TOPFAN                = 30
	SVC_SNSMESSAGE            = 31
	SVC_SNSMODE               = 32
	SVC_SENDBALLOONSUB        = 33
	SVC_SENDFANLETRTRERSUB    = 34
	SVC_TOPFANSUB             = 35
	SVC_BJSTICKERITEM         = 36
	SVC_CHOCOLATE             = 37
	SVC_CHOCOLATESUB          = 38
	SVC_TOPCLAN               = 39
	SVC_TOPCLANSUB            = 40
	SVC_SUPERCHAT             = 41
	SVC_UPDATETICKET          = 42
	SVC_NOTIGAMERANKER        = 43
	SVC_STARCOIN              = 44
	SVC_SENDQUICKVIEW         = 45
	SVC_ITEMSTATUS            = 46
	SVC_ITEMUSING             = 47
	SVC_USEQUICKVIEW          = 48
	SVC_NOTIFY_POLL           = 50
	SVC_CHATBLOCKMODE         = 51
	SVC_BDM_ADDBLACKINFO      = 52
	SVC_SETBROADINFO          = 53
	SVC_BAN_WORD              = 54
	SVC_SENDADMINNOTICE       = 58
	SVC_FREECAT_OWNER_JOIN    = 65
	SVC_BUYGOODS              = 70
	SVC_BUYGOODSSUB           = 71
	SVC_SENDPROMOTION         = 72
	SVC_NOTIFY_VR             = 74
	SVC_NOTIFY_MOBBROAD_PAUSE = 75
	SVC_KICK_AND_CANCEL       = 76
	SVC_KICK_USERLIST         = 77
	SVC_ADMIN_CHUSER          = 78
	SVC_CLIDOBAEINFO          = 79
	SVC_VOD_BALLOON           = 86
	SVC_ADCON_EFFECT          = 87
	SVC_SVC_KICK_MSG_STATE    = 90
	SVC_FOLLOW_ITEM           = 91 // "\x1b\t009100004500\f2884\fmygomiee\fzlah0864\f종우패는호종\f7\f"
	SVC_ITEM_SELL_EFFECT      = 92
	SVC_FOLLOW_ITEM_EFFECT    = 93
	SVC_TRANSLATION_STATE     = 94
	SVC_TRANSLATION           = 95
	SVC_GIFT_TICKET           = 102
	SVC_VODADCON              = 103
	SVC_BJ_NOTICE             = 104
	SVC_VIDEOBALLOON          = 105
	SVC_STATION_ADCON         = 107 /* "\x1b\t010700018800\fmygomiee\fzzh9988\fJ.Render\f1\f//res.afreecatv.com/new_player/items/adballoon/ceremony/pc_1.png?t=1642055787\f방송국에서 J.Render님이
	애드벌룬 1개를 선물 하셨습니다.\f2884\f" */
	SVC_SENDSUBSCRIPTION  = 108
	SVC_OGQ_EMOTICON      = 109
	SVC_ITEM_DROPS        = 111
	SVC_VIDEOBALLOON_LINK = 117
	SVC_OGQ_EMOTICON_GIFT = 118
	SVC_AD_IN_BROAD_JSON  = 119
	SVC_GEM_ITEMSEND      = 120
	SVC_MISSION           = 121 /* 도전 미션 ? */
	SVC_LIVE_CAPTION      = 122
	SVC_MISSION_SETTLE    = 125
	SVC_SET_ADMIN_FLAG    = 126
)

/*
"\x1b\t001800006200\fmygomiee\fstdeviler\f쟌다링\f10\f0\f0\f2884\f10\f0\f0\fkor_custom12\f"
{User:{ID:stdeviler Name:쟌다링 SubscribeMonth:0 Flag:{Flag1:{Admin:false Hidden:false BJ:false Dumb:false Guest:false Fanclub:false AutoManager:false ManagerList:false Manager:false Female:false AutoDumb:false DumbBlind:false DobaeBlind:false DobaeBlind2:false ExitUser:false Mobile:false TopFan:false Realname:false NoDirect:false GlobalApp:false QuickView:false SptrSticker:false Chromecast:false Follower:false NotiVodBalloon:false NotiTopFan:false} Flag2:{GlobalPC:false Clan:false TopClan:false Top20:false GameGod:false ATagAllow:false NoSuperChat:false NoRecvChat:false Flash:false LGGame:false Employee:false CleanAti:false Police:false AdminChat:false PC:false Specify:false}}} Count:10}
"\x1b\t000500014800\f이 의견은 마이곰이 개인적인 의견이며 저희랑은 상관이 없음을 알려드립니다\fstdeviler\f0\f0\f3\f쟌다링\f65568|163840\f-1\f"
"\x1b\t001800005700\fmygomiee\fneny9965\fPnec4\f33\f0\f0\f2884\f33\f0\f0\fkor_custom14\f"
{User:{ID:neny9965 Name:Pnec4 SubscribeMonth:0 Flag:{Flag1:{Admin:false Hidden:false BJ:false Dumb:false Guest:false Fanclub:false AutoManager:false ManagerList:false Manager:false Female:false AutoDumb:false DumbBlind:false DobaeBlind:false DobaeBlind2:false ExitUser:false Mobile:false TopFan:false Realname:false NoDirect:false GlobalApp:false QuickView:false SptrSticker:false Chromecast:false Follower:false NotiVodBalloon:false NotiTopFan:false} Flag2:{GlobalPC:false Clan:false TopClan:false Top20:false GameGod:false ATagAllow:false NoSuperChat:false NoRecvChat:false Flash:false LGGame:false Employee:false CleanAti:false Police:false AdminChat:false PC:false Specify:false}}} Count:33}
"\x1b\t000500008500\f마이곰이님 개인의 의견입니다,,\fneny9965\f0\f0\f3\fPnec4\f268501024|163840\f1\f"
*/
