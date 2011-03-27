package main
 
// Includes functions for sending data to the server.
 
import (
	"net"
)

func SendHeader(c net.Conn, val byte) {
	WriteByte(c, val)
}

func SendChat(c net.Conn, message string) {
	SendHeader(c, 0x03)
	if len(message) > 100 {
		log0.Println("Chat message too long! Truncated to 100 characters.")
		message = message[:100]
	}
	WriteString(c, message)
}

func SendRespawn(c net.Conn) {
	SendHeader(c, 0x09)
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

func SendPlayerPos(c net.Conn, pp *PlayerCoords) {
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

func SendPlayerDigging(c net.Conn, bc *BlockCoords, status, face byte) {
	SendHeader(c, 0x0E)
	WriteByte(c, status)
	WriteInt(c, bc.X)
	WriteByte(c, byte(bc.Y))
	WriteInt(c, bc.Z)
	WriteByte(c, face)
}