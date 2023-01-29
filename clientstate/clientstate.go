package clientstate

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type (
	Platform     uint32
	Product      uint32
	ProtocolType byte
)

const (
	PLATFORM_IX86 Platform = 0x49583836 // Windows (x86)
	PLATFORM_PMAC Platform = 0x504D4143 // macOS (PowerPC)
	PLATFORM_XMAC Platform = 0x584D4143 // macOS (x86)
	PLATFORM_ZERO Platform = 0x00000000 // Null (Zero)
)

const (
	PRODUCT_CHAT Product = 0x43484154 // Chat Gateway Client
	PRODUCT_D2DV Product = 0x44324456 // Diablo II
	PRODUCT_D2XP Product = 0x44325850 // Diablo II Lord of Destruction
	PRODUCT_DRTL Product = 0x4452544C // Diablo Retail
	PRODUCT_DSHR Product = 0x44534852 // Diablo Shareware
	PRODUCT_JSTR Product = 0x4A535452 // Starcraft Japanese
	PRODUCT_SEXP Product = 0x53455850 // Starcraft Broodwar
	PRODUCT_SSHR Product = 0x53534852 // Starcraft Shareware
	PRODUCT_STAR Product = 0x53544152 // Starcraft Original
	PRODUCT_W2BN Product = 0x5732424E // Warcraft II Battle.net Edition
	PRODUCT_W3DM Product = 0x5733444D // Warcraft III Demo
	PRODUCT_W3XP Product = 0x57335850 // Warcraft III The Frozen Throne
	PRODUCT_WAR3 Product = 0x57415233 // Warcraft III Reign of Chaos
	PRODUCT_ZERO Product = 0x00000000 // Null (Zero)
)

type ClientState struct {
	ClientLocalIP        uint32
	Conn                 net.Conn
	CountryCode          []byte
	CountryCodeAbbr      []byte
	CountryName          []byte
	CountryNameAbbr      []byte
	LocaleSystemLCID     uint32
	LocaleUserLanguageId uint32
	LocaleUserLCID       uint32
	Ping                 int32
	PingCookie           uint32
	Platform             Platform
	Product              Product
	ProtocolType         ProtocolType
	RemoteAddr           net.Addr
	TimezoneBias         int32
	Username             []byte
	VersionId            uint32 // also known as "version byte" in other software
}

var clientStates = sync.Map{}

var platformNames = map[Platform]string{
	PLATFORM_IX86: "Windows (x86)",
	PLATFORM_PMAC: "macOS (PowerPC)",
	PLATFORM_XMAC: "macOS (x86)",
	PLATFORM_ZERO: "(null)",
}

var productNames = map[Product]string{
	PRODUCT_CHAT: "Chat Gateway Client",
	PRODUCT_D2DV: "Diablo II",
	PRODUCT_D2XP: "Diablo II Lord of Destruction",
	PRODUCT_DRTL: "Diablo",
	PRODUCT_DSHR: "Diablo Shareware",
	PRODUCT_JSTR: "Starcraft Japan",
	PRODUCT_SEXP: "Starcraft Broodwar",
	PRODUCT_SSHR: "Starcraft Shareware",
	PRODUCT_STAR: "Starcraft",
	PRODUCT_W2BN: "Warcraft II BNE",
	PRODUCT_W3DM: "Warcraft III Demo",
	PRODUCT_W3XP: "Warcraft III The Frozen Throne",
	PRODUCT_WAR3: "Warcraft III Reign of Chaos",
	PRODUCT_ZERO: "(null)",
}

func AddClientState(conn net.Conn, state *ClientState) {
	clientStates.Store(conn, state)
}

func RemoveClientState(conn net.Conn) {
	clientStates.Delete(conn)
}

func GetClientState(conn net.Conn) (*ClientState, bool) {
	state, ok := clientStates.Load(conn)
	if !ok {
		return nil, false
	}
	return state.(*ClientState), true
}

func PlatformToName(value Platform) string {
	if name, ok := platformNames[value]; ok {
		return name
	}
	return fmt.Sprintf("Unknown (%08X)", value)
}

func ProductToName(value Product) string {
	if name, ok := productNames[value]; ok {
		return name
	}
	return fmt.Sprintf("Unknown (%08X)", value)
}

func ReadProtocolType(conn io.Reader) (ProtocolType, error) {
	var buf [1]byte
	_, err := io.ReadFull(conn, buf[:])
	if err != nil {
		return 0, err
	}
	return ProtocolType(buf[0]), nil
}
