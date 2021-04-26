package handler

import (
	"Stowaway/global"
	"Stowaway/protocol"
)

func LetOffline(route string, uuid string) {
	sMessage := protocol.PrepareAndDecideWhichSProtoToLower(global.G_Component.Conn, global.G_Component.Secret, global.G_Component.UUID)

	header := &protocol.Header{
		Sender:      protocol.ADMIN_UUID,
		Accepter:    uuid,
		MessageType: protocol.OFFLINE,
		RouteLen:    uint32(len([]byte(route))),
		Route:       route,
	}

	offlineMess := &protocol.Offline{
		OK: 1,
	}

	protocol.ConstructMessage(sMessage, header, offlineMess, false)
	sMessage.SendMessage()
}
