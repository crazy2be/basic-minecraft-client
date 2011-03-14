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
		case 9:
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
							fmt.Println("Second argument should be 'x', 'y', or 'z', followed by the amount")
					}
					SetPlayerPos(c, pp)
				case "sleep":
					amount := 0.0
					fmt.Scan(&amount)
					time.Sleep(int64(amount*1000*1000*1000))
				case "say", "chat":
					message := ""
					fmt.Scanf("%q", &message)
					SendChat(c, message)
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
	for i := 0; i < 100000; i++ {
		NextPacket(c)
	}
}

var userName *string

func main() {
	userName = flag.String("username", "test", "The username that this client should join with.")
	flag.Parse()
	InitLogs()
	c, err := net.Dial("tcp", "", "68.144.126.229:25565")
	//24.13.132.130:25565
	if err != nil {
		log.Fatal(err)
	}
	login(c)
	
}