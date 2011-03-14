package main
 
// Includes functions for sending data to the server.
 
import (
	"net"
)

func SendHeader(c net.Conn, val byte) {
	WriteByte(c, val)
}

func SendPlayerPositionAndLook(c net.Conn, x, stance, y, z float64, yaw, pitch float32, onGround byte) {
	SendHeader(c, 0x0D)
	WriteDouble(c, x)
	WriteDouble(c, stance)
	WriteDouble(c, y)
	WriteDouble(c, z)
	WriteFloat(c, yaw)
	WriteFloat(c, pitch)
	WriteByte(c, onGround)
}

func SendPlayerPos(c net.Conn, pp *PlayerPos) {
	SendHeader(c, 0x0B)
	WriteDouble(c, pp.X)
	WriteDouble(c, pp.Y)
	WriteDouble(c, pp.YTop)
	WriteDouble(c, pp.Z)
	var ogbyte byte
	if pp.OnGround == false {
		ogbyte = 0
	} else {
		ogbyte = 1
	}
	WriteByte(c, ogbyte)
}

func SendChat(c net.Conn, message string) {
	SendHeader(c, 0x03)
	WriteString(c, message)
}