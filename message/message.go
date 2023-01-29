package message

import (
	"encoding/binary"
	"fmt"
	"io"
)

type MessageId byte

type Message struct {
	ID     MessageId
	Length uint16
	Body   []byte
}

const (
	SID_ANNOUNCEMENT             MessageId = 0x20
	SID_AUTH_ACCOUNTCHANGE       MessageId = 0x55
	SID_AUTH_ACCOUNTCHANGEPROOF  MessageId = 0x56
	SID_AUTH_ACCOUNTCREATE       MessageId = 0x52
	SID_AUTH_ACCOUNTLOGON        MessageId = 0x53
	SID_AUTH_ACCOUNTLOGONPROOF   MessageId = 0x54
	SID_AUTH_ACCOUNTUPGRADE      MessageId = 0x57
	SID_AUTH_ACCOUNTUPGRADEPROOF MessageId = 0x58
	SID_AUTH_CHECK               MessageId = 0x51
	SID_AUTH_INFO                MessageId = 0x50
	SID_CDKEY                    MessageId = 0x30
	SID_CDKEY2                   MessageId = 0x36
	SID_CDKEY3                   MessageId = 0x42
	SID_CHANGEEMAIL              MessageId = 0x5B
	SID_CHANGEPASSWORD           MessageId = 0x31
	SID_CHATCOMMAND              MessageId = 0x0E
	SID_CHATEVENT                MessageId = 0x0F
	SID_CHECKAD                  MessageId = 0x15
	SID_CHECKDATAFILE            MessageId = 0x32
	SID_CHECKDATAFILE2           MessageId = 0x3C
	SID_CLANCREATIONINVITATION   MessageId = 0x72
	SID_CLANDISBAND              MessageId = 0x73
	SID_CLANFINDCANDIDATES       MessageId = 0x70
	SID_CLANINFO                 MessageId = 0x75
	SID_CLANINVITATION           MessageId = 0x77
	SID_CLANINVITATIONRESPONSE   MessageId = 0x79
	SID_CLANINVITEMULTIPLE       MessageId = 0x71
	SID_CLANMAKECHIEFTAIN        MessageId = 0x74
	SID_CLANMEMBERINFORMATION    MessageId = 0x82
	SID_CLANMEMBERLIST           MessageId = 0x7D
	SID_CLANMEMBERRANKCHANGE     MessageId = 0x81
	SID_CLANMEMBERREMOVED        MessageId = 0x7E
	SID_CLANMEMBERSTATUSCHANGE   MessageId = 0x7F
	SID_CLANMOTD                 MessageId = 0x7C
	SID_CLANQUITNOTIFY           MessageId = 0x76
	SID_CLANRANKCHANGE           MessageId = 0x7A
	SID_CLANREMOVEMEMBER         MessageId = 0x78
	SID_CLANSETMOTD              MessageId = 0x7B
	SID_CLICKAD                  MessageId = 0x16
	SID_CLIENTID                 MessageId = 0x05
	SID_CLIENTID2                MessageId = 0x1E
	SID_CREATEACCOUNT            MessageId = 0x2A
	SID_CREATEACCOUNT2           MessageId = 0x3D
	SID_DISPLAYAD                MessageId = 0x21
	SID_ENTERCHAT                MessageId = 0x0A
	SID_EXTRAWORK                MessageId = 0x4B
	SID_FINDLADDERUSER           MessageId = 0x2F
	SID_FLOODDETECTED            MessageId = 0x13
	SID_FRIENDSADD               MessageId = 0x67
	SID_FRIENDSLIST              MessageId = 0x65
	SID_FRIENDSPOSITION          MessageId = 0x69
	SID_FRIENDSREMOVE            MessageId = 0x68
	SID_FRIENDSUPDATE            MessageId = 0x66
	SID_GAMEDATAADDRESS          MessageId = 0x1B
	SID_GAMEPLAYERSEARCH         MessageId = 0x60
	SID_GAMERESULT               MessageId = 0x2C
	SID_GETADVLISTEX             MessageId = 0x09
	SID_GETCHANNELLIST           MessageId = 0x0B
	SID_GETFILETIME              MessageId = 0x33
	SID_GETICONDATA              MessageId = 0x2D
	SID_GETLADDERDATA            MessageId = 0x2E
	SID_JOINCHANNEL              MessageId = 0x0C
	SID_LEAVECHAT                MessageId = 0x10
	SID_LEAVEGAME                MessageId = 0x1F
	SID_LOCALEINFO               MessageId = 0x12
	SID_LOGONCHALLENGE           MessageId = 0x28
	SID_LOGONCHALLENGEEX         MessageId = 0x1D
	SID_LOGONREALMEX             MessageId = 0x3E
	SID_LOGONRESPONSE            MessageId = 0x29
	SID_LOGONRESPONSE2           MessageId = 0x3A
	SID_MESSAGEBOX               MessageId = 0x19
	SID_NETGAMEPORT              MessageId = 0x45
	SID_NEWS_INFO                MessageId = 0x46
	SID_NOTIFYJOIN               MessageId = 0x22
	SID_NULL                     MessageId = 0x00
	SID_OPTIONALWORK             MessageId = 0x4A
	SID_PING                     MessageId = 0x25
	SID_PROFILE                  MessageId = 0x35
	SID_QUERYADURL               MessageId = 0x41
	SID_QUERYREALMS              MessageId = 0x34
	SID_QUERYREALMS2             MessageId = 0x40
	SID_READCOOKIE               MessageId = 0x24
	SID_READMEMORY               MessageId = 0x17
	SID_READUSERDATA             MessageId = 0x26
	SID_REGISTRY                 MessageId = 0x18
	SID_REPORTCRASH              MessageId = 0x5D
	SID_REPORTVERSION            MessageId = 0x07
	SID_REQUIREDWORK             MessageId = 0x4C
	SID_RESETPASSWORD            MessageId = 0x5A
	SID_SERVERLIST               MessageId = 0x04
	SID_SETEMAIL                 MessageId = 0x59
	SID_STARTADVEX               MessageId = 0x08
	SID_STARTADVEX2              MessageId = 0x1A
	SID_STARTADVEX3              MessageId = 0x1C
	SID_STARTVERSIONING          MessageId = 0x06
	SID_STARTVERSIONING2         MessageId = 0x3F
	SID_STOPADV                  MessageId = 0x02
	SID_SWITCHPRODUCT            MessageId = 0x5C
	SID_SYSTEMINFO               MessageId = 0x2B
	SID_TOURNAMENT               MessageId = 0x4E
	SID_UDPPINGRESPONSE          MessageId = 0x14
	SID_WARCRAFTGENERAL          MessageId = 0x44
	SID_WARCRAFTUNKNOWN          MessageId = 0x43
	SID_WARDEN                   MessageId = 0x5E
	SID_WRITECOOKIE              MessageId = 0x23
	SID_WRITEUSERDATA            MessageId = 0x27
)

var messageIdNames = map[MessageId]string{
	SID_ANNOUNCEMENT:             "SID_ANNOUNCEMENT",
	SID_AUTH_ACCOUNTCHANGE:       "SID_AUTH_ACCOUNTCHANGE",
	SID_AUTH_ACCOUNTCHANGEPROOF:  "SID_AUTH_ACCOUNTCHANGEPROOF",
	SID_AUTH_ACCOUNTCREATE:       "SID_AUTH_ACCOUNTCREATE",
	SID_AUTH_ACCOUNTLOGON:        "SID_AUTH_ACCOUNTLOGON",
	SID_AUTH_ACCOUNTLOGONPROOF:   "SID_AUTH_ACCOUNTLOGONPROOF",
	SID_AUTH_ACCOUNTUPGRADE:      "SID_AUTH_ACCOUNTUPGRADE",
	SID_AUTH_ACCOUNTUPGRADEPROOF: "SID_AUTH_ACCOUNTUPGRADEPROOF",
	SID_AUTH_CHECK:               "SID_AUTH_CHECK",
	SID_AUTH_INFO:                "SID_AUTH_INFO",
	SID_CDKEY:                    "SID_CDKEY",
	SID_CDKEY2:                   "SID_CDKEY2",
	SID_CDKEY3:                   "SID_CDKEY3",
	SID_CHANGEEMAIL:              "SID_CHANGEEMAIL",
	SID_CHANGEPASSWORD:           "SID_CHANGEPASSWORD",
	SID_CHATCOMMAND:              "SID_CHATCOMMAND",
	SID_CHATEVENT:                "SID_CHATEVENT",
	SID_CHECKAD:                  "SID_CHECKAD",
	SID_CHECKDATAFILE:            "SID_CHECKDATAFILE",
	SID_CHECKDATAFILE2:           "SID_CHECKDATAFILE2",
	SID_CLANCREATIONINVITATION:   "SID_CLANCREATIONINVITATION",
	SID_CLANDISBAND:              "SID_CLANDISBAND",
	SID_CLANFINDCANDIDATES:       "SID_CLANFINDCANDIDATES",
	SID_CLANINFO:                 "SID_CLANINFO",
	SID_CLANINVITATION:           "SID_CLANINVITATION",
	SID_CLANINVITATIONRESPONSE:   "SID_CLANINVITATIONRESPONSE",
	SID_CLANINVITEMULTIPLE:       "SID_CLANINVITEMULTIPLE",
	SID_CLANMAKECHIEFTAIN:        "SID_CLANMAKECHIEFTAIN",
	SID_CLANMEMBERINFORMATION:    "SID_CLANMEMBERINFORMATION",
	SID_CLANMEMBERLIST:           "SID_CLANMEMBERLIST",
	SID_CLANMEMBERRANKCHANGE:     "SID_CLANMEMBERRANKCHANGE",
	SID_CLANMEMBERREMOVED:        "SID_CLANMEMBERREMOVED",
	SID_CLANMEMBERSTATUSCHANGE:   "SID_CLANMEMBERSTATUSCHANGE",
	SID_CLANMOTD:                 "SID_CLANMOTD",
	SID_CLANQUITNOTIFY:           "SID_CLANQUITNOTIFY",
	SID_CLANRANKCHANGE:           "SID_CLANRANKCHANGE",
	SID_CLANREMOVEMEMBER:         "SID_CLANREMOVEMEMBER",
	SID_CLANSETMOTD:              "SID_CLANSETMOTD",
	SID_CLICKAD:                  "SID_CLICKAD",
	SID_CLIENTID:                 "SID_CLIENTID",
	SID_CLIENTID2:                "SID_CLIENTID2",
	SID_CREATEACCOUNT:            "SID_CREATEACCOUNT",
	SID_CREATEACCOUNT2:           "SID_CREATEACCOUNT2",
	SID_DISPLAYAD:                "SID_DISPLAYAD",
	SID_ENTERCHAT:                "SID_ENTERCHAT",
	SID_EXTRAWORK:                "SID_EXTRAWORK",
	SID_FINDLADDERUSER:           "SID_FINDLADDERUSER",
	SID_FLOODDETECTED:            "SID_FLOODDETECTED",
	SID_FRIENDSADD:               "SID_FRIENDSADD",
	SID_FRIENDSLIST:              "SID_FRIENDSLIST",
	SID_FRIENDSPOSITION:          "SID_FRIENDSPOSITION",
	SID_FRIENDSREMOVE:            "SID_FRIENDSREMOVE",
	SID_FRIENDSUPDATE:            "SID_FRIENDSUPDATE",
	SID_GAMEDATAADDRESS:          "SID_GAMEDATAADDRESS",
	SID_GAMEPLAYERSEARCH:         "SID_GAMEPLAYERSEARCH",
	SID_GAMERESULT:               "SID_GAMERESULT",
	SID_GETADVLISTEX:             "SID_GETADVLISTEX",
	SID_GETCHANNELLIST:           "SID_GETCHANNELLIST",
	SID_GETFILETIME:              "SID_GETFILETIME",
	SID_GETICONDATA:              "SID_GETICONDATA",
	SID_GETLADDERDATA:            "SID_GETLADDERDATA",
	SID_JOINCHANNEL:              "SID_JOINCHANNEL",
	SID_LEAVECHAT:                "SID_LEAVECHAT",
	SID_LEAVEGAME:                "SID_LEAVEGAME",
	SID_LOCALEINFO:               "SID_LOCALEINFO",
	SID_LOGONCHALLENGE:           "SID_LOGONCHALLENGE",
	SID_LOGONCHALLENGEEX:         "SID_LOGONCHALLENGEEX",
	SID_LOGONREALMEX:             "SID_LOGONREALMEX",
	SID_LOGONRESPONSE:            "SID_LOGONRESPONSE",
	SID_LOGONRESPONSE2:           "SID_LOGONRESPONSE2",
	SID_MESSAGEBOX:               "SID_MESSAGEBOX",
	SID_NETGAMEPORT:              "SID_NETGAMEPORT",
	SID_NEWS_INFO:                "SID_NEWS_INFO",
	SID_NOTIFYJOIN:               "SID_NOTIFYJOIN",
	SID_NULL:                     "SID_NULL",
	SID_OPTIONALWORK:             "SID_OPTIONALWORK",
	SID_PING:                     "SID_PING",
	SID_PROFILE:                  "SID_PROFILE",
	SID_QUERYADURL:               "SID_QUERYADURL",
	SID_QUERYREALMS:              "SID_QUERYREALMS",
	SID_QUERYREALMS2:             "SID_QUERYREALMS2",
	SID_READCOOKIE:               "SID_READCOOKIE",
	SID_READMEMORY:               "SID_READMEMORY",
	SID_READUSERDATA:             "SID_READUSERDATA",
	SID_REGISTRY:                 "SID_REGISTRY",
	SID_REPORTCRASH:              "SID_REPORTCRASH",
	SID_REPORTVERSION:            "SID_REPORTVERSION",
	SID_REQUIREDWORK:             "SID_REQUIREDWORK",
	SID_RESETPASSWORD:            "SID_RESETPASSWORD",
	SID_SERVERLIST:               "SID_SERVERLIST",
	SID_SETEMAIL:                 "SID_SETEMAIL",
	SID_STARTADVEX:               "SID_STARTADVEX",
	SID_STARTADVEX2:              "SID_STARTADVEX2",
	SID_STARTADVEX3:              "SID_STARTADVEX3",
	SID_STARTVERSIONING:          "SID_STARTVERSIONING",
	SID_STARTVERSIONING2:         "SID_STARTVERSIONING2",
	SID_STOPADV:                  "SID_STOPADV",
	SID_SWITCHPRODUCT:            "SID_SWITCHPRODUCT",
	SID_SYSTEMINFO:               "SID_SYSTEMINFO",
	SID_TOURNAMENT:               "SID_TOURNAMENT",
	SID_UDPPINGRESPONSE:          "SID_UDPPINGRESPONSE",
	SID_WARCRAFTGENERAL:          "SID_WARCRAFTGENERAL",
	SID_WARCRAFTUNKNOWN:          "SID_WARCRAFTUNKNOWN",
	SID_WARDEN:                   "SID_WARDEN",
	SID_WRITECOOKIE:              "SID_WRITECOOKIE",
	SID_WRITEUSERDATA:            "SID_WRITEUSERDATA",
}

func MessageIdToName(id MessageId) string {
	if name, ok := messageIdNames[id]; ok {
		return name
	}
	return fmt.Sprintf("SID_UNKNOWN_%02X", id)
}

func ReadMessage(conn io.Reader) (*Message, error) {
	var header [4]byte
	_, err := io.ReadFull(conn, header[:])
	if err != nil {
		return nil, err
	}

	if header[0] != 0xFF {
		return nil, fmt.Errorf("invalid message header")
	}

	length := binary.LittleEndian.Uint16(header[2:4])
	body := make([]byte, length-4)
	_, err = io.ReadFull(conn, body)
	if err != nil {
		return nil, err
	}

	return &Message{
		ID:     MessageId(header[1]),
		Length: length,
		Body:   body,
	}, nil
}

func ReadProtocolByte(conn io.Reader) (byte, error) {
	var buf [1]byte
	_, err := io.ReadFull(conn, buf[:])
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}
