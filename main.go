package main

import (
	"net"
	"log"
	"time"
	"flag"
	"fmt"
	//"io"
	"os"
)

func NextPacket(c net.Conn) {
	b, err := ReadByte(c)
	if err != nil {
		log.Println("Error reading next packet:", err)
		os.Exit(1)
	}
	switch (b) {
		case 0x00:
			// Keep-alive ping
			log.Println("Received Keep-alive")
		case 1:
			HandleLogin(c)
		case 2:
			HandleHandshake(c)
		case 3:
			HandleChat(c)
		case 4:
			HandleTime(c)
		case 5:
			HandleEquipment(c)
		case 6:
			HandleSpawnPos(c)
		case 0x07:
			HandleUseEntity(c)
		case 8:
			HandleHealth(c)
		case 0x09:
			HandleRespawn(c)
		case 0x0D:
			HandlePlayerPositionAndLook(c)
		//case 0x11:
		//	HandleBed(c)
		case 0x12:
			HandleAnimation(c)
		case 0x14:
			HandleNamedEntitySpawn(c)
		case 0x15:
			HandlePickupSpawn(c)
		case 0x16:
			HandleCollectItem(c)
		case 0x17:
			HandleAddObject(c)
		case 0x18:
			HandleMobSpawn(c)
		case 0x19:
			HandlePainting(c)
		case 0x1C:
			HandleEntityVelocity(c)
		case 0x1D:
			HandleDestroyEntity(c)
		case 0x1E:
			HandleEntity(c)
		case 0x1F:
			HandleEntityRelativeMove(c)
		case 0x20:
			HandleEntityLook(c)
		case 0x21:
			HandleEntityLookAndRelativeMove(c)
		case 0x22:
			HandleEntityTeleport(c)
		case 0x26:
			HandleEntityStatus(c)
		case 0x28:
			HandleEntityMetadata(c)
		case 0x32:
			HandlePreChunk(c)
		case 0x33:
			HandleChunk(c)
		case 0x34:
			HandleMultiBlockChange(c)
		case 0x35:
			HandleBlockChange(c)
		case 0x36:
			HandlePlayNoteBlock(c)
		case 0x3C:
			HandleExplosion(c)
		case 0x67:
			HandleSetSlot(c)
		case 0x68:
			HandleWindowItems(c)
		case 0x82:
			HandleSignUpdate(c)
		case 0xFF:
			HandleKick(c)
		default:
			log.Println("Unknown packet", b, "! Trouble Ahead!")
			os.Exit(0)
	}
}

func KeepAlive(c net.Conn) {
	for {
		// 10 seconds
		time.Sleep(10*1000*1000*1000)
		
		WriteByte(c, 0)
	}
}

// TODO: Hackish, should be in coords and provide a mechanism to subtract two coord sets.
func GetPlayerOffset(cmd string) *PlayerCoords {
	dir := ""
	fmt.Scan(&dir)
	amount := 0.0
	fmt.Scan(&amount)
	pp := GetPlayerPos()
	switch (dir) {
		case "x":
			pp.X += amount
		case "y":
			pp.Y += amount
		case "z":
			pp.Z += amount
		default:
			fmt.Printf("Syntax: %s <dir> <amount>\nWhere %s is one of [x, y, z], and amount is a floating point number, relative to the current player position.\n", cmd, cmd)
	}
	return pp
}

func HandleCommands(c net.Conn) {
	for {
		cmd := ""
		n, err := fmt.Scan(&cmd)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		if n == 1 {
			switch (cmd) {
				case "quit", "exit", "stop":
					os.Exit(0)
				case "pos":
					fmt.Println("Player position:", GetPlayerPos())
				case "move":
					pp := GetPlayerOffset("move")
					SetPlayerPos(c, pp)
				case "look":
					pp := GetPlayerOffset("look")
					bc := pp.BlockCoords()
					b := GetBlock(bc)
					if b == nil {
						fmt.Println("Block at", bc, pp, "not loaded!")
						continue
					}
					fmt.Println("Block at", bc, ":", b)
				case "dig":
					pp := GetPlayerOffset("dig")
					bc := pp.BlockCoords()
					// Do we need to detect this?
// 					face := uint8(0)
// 					switch (dir) {
// 						case "x":
// 							if amount > 0 {
// 								face = 5
// 							} else {
// 								face = 4
// 							}
// 						case "y":
// 							if amount > 0 {
// 								face = 1
// 							} else {
// 								face = 0
// 							}
// 						case "z":
// 							if amount > 0 {
// 								face = 3
// 							} else {
// 								face = 2
// 							}
// 					}
					fmt.Println("Started digging block at", bc, pp)
					SendPlayerDigging(c, bc, 0, 0)
					// TODO: Figure out actual time requirements and use those based on block type. For now, this works.
					time.Sleep(2*1000*1000*1000)
					SendPlayerDigging(c, bc, 2, 0)
				case "sleep":
					amount := 0.0
					fmt.Scan(&amount)
					time.Sleep(int64(amount*1000*1000*1000))
				case "say", "chat":
					message := ""
					fmt.Scanf("%q", &message)
					SendChat(c, message)
				case "respawn":
					SendRespawn(c)
				default:
					log.Println("Unrecognized command", cmd)
			}
		}
	}
}

func login(c net.Conn) {
	// First, do the handshake (0x02)
	WriteByte(c, 2)
	WriteString(c, *userName)
	
	NextPacket(c)
	
	// Then, send the login request (0x01)
	WriteByte(c, 1)
	WriteInt(c, 9)
	WriteString(c, *userName)
	WriteString(c, "pass")
	WriteLong(c, 0)
	WriteByte(c, 0)
	
	go KeepAlive(c)
	go func() {
		HandleCommands(c)
	}()
	for {
		NextPacket(c)
	}
}

var userName *string

func main() {
	userName = flag.String("username", "test", "The username that this client should join with.")
	serverIP := flag.String("serverip", "68.144.126.229:25565", "The server to join, including the port number.")
	flag.Parse()
	InitLogs()
	c, err := net.Dial("tcp", "", *serverIP)
	//24.13.132.130:25565
	//68.144.126.229:25565
	//68.147.253.198:25565
	if err != nil {
		log.Fatal(err)
	}
	login(c)
	
}
