package main

import (
	"net"
	"log"
	"os"
	"bytes"
	"compress/flate"
)

var nlog *log.Logger

func init() {
	nfile, err := os.Open(os.DevNull, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("WTF. Could not open null device!")
	}
	nlog = log.New(nfile, "", 0)
}

func HandleLogin(c net.Conn) {
	eid, _ := ReadInt(c)
	u1, _ := ReadString(c)
	u2, _ := ReadString(c)
	seed, _ := ReadLong(c)
	dimension, _ := ReadByte(c)
	log1.Println("Logged in! ID:", eid, "other stuff:", u1, u2, seed, dimension)
}

func HandleHandshake(c net.Conn) {
	hash, _ := ReadString(c)
	log1.Println("Got login hash:", hash)
}

func HandleChat(c net.Conn) {
	chat, _ := ReadString(c)
	log1.Println("Got chat message:", chat)
}

func HandleTime(c net.Conn) {
	time, _ := ReadLong(c)
	DataWorldTime = time
	log7.Println("Got time update:", time)
}

func HandleEquipment(c net.Conn) {
	eid, _ := ReadInt(c)
	slot, _ := ReadShort(c)
	iid, _ := ReadShort(c)
	u1, _ := ReadShort(c)
	log6.Println("Got handle equipment packet.", eid, slot, iid, u1)
}

func HandleSpawnPos(c net.Conn) {
	x, _ := ReadInt(c)
	y, _ := ReadInt(c)
	z, _ := ReadInt(c)
	log3.Println("Spawn position is:", x, y, z)
}

func HandleUseEntity(c net.Conn) {
	eid, _ := ReadInt(c)
	target, _ := ReadInt(c)
	click, _ := ReadByte(c)
	log3.Println("Use Entity:", eid, target, click)
}

func HandleHealth(c net.Conn) {
	health, _ := ReadShort(c)
	log2.Println("Player health is:", health)
}

func HandleRespawn(c net.Conn) {
	log2.Println("Received Respawn packet")
}

func HandlePlayerPositionAndLook(c net.Conn) {
	x, _ := ReadDouble(c)
	y, _ := ReadDouble(c)
	stance, _ := ReadDouble(c)
	z, _ := ReadDouble(c)
	yaw, _ := ReadFloat(c)
	pitch, _ := ReadFloat(c)
	onGroundraw, _ := ReadByte(c)
	var onGround bool
	if onGroundraw != 0 {
		onGround = true
	}
	SendPlayerPositionAndLook(c, x, stance, y, z, yaw, pitch, onGroundraw)
	UpdatePlayerPos(&PlayerPos{x, y, stance, z, onGround})
	UpdatePlayerLook(&PlayerLook{yaw, pitch, onGround})
	log4.Println("Player Position & Look Packet", x, y, stance, z, yaw, pitch, onGround)
}

func HandleAnimation(c net.Conn) {
	eid, _ := ReadInt(c)
	animation, _ := ReadByte(c)
	log4.Println("Animation:", eid, animation)
}

func HandleBed(c net.Conn) {
	
}

func HandleNamedEntitySpawn(c net.Conn) {
	eid, _ := ReadInt(c)
	name, _ := ReadString(c)
	x, _ := ReadInt(c)
	y, _ := ReadInt(c)
	z, _ := ReadInt(c)
	rot, _ := ReadByte(c)
	pitch, _ := ReadByte(c)
	item, _ := ReadShort(c)
	log7.Println("Named Entity Spawn:", eid, name, x, y, z, rot, pitch, item)
}

func HandlePickupSpawn(c net.Conn) {
	eid, _ := ReadInt(c)
	item, _ := ReadShort(c)
	count, _ := ReadByte(c)
	data, _ := ReadShort(c)
	x, _ := ReadInt(c)
	y, _ := ReadInt(c)
	z, _ := ReadInt(c)
	rot, _ := ReadByte(c)
	pitch, _ := ReadByte(c)
	roll, _ := ReadByte(c)
	log8.Println("Pickup spawn:", eid, item, count, data, x, y, z, rot, pitch, roll)
}

func HandleCollectItem(c net.Conn) {
	collected, _ := ReadInt(c)
	collector, _ := ReadInt(c)
	log4.Println("Collected Item:", collected, collector)
}

func HandleAddObject(c net.Conn) {
	eid, _ := ReadInt(c)
	typ, _ := ReadByte(c)
	x, _ := ReadInt(c)
	y, _ := ReadInt(c)
	z, _ := ReadInt(c)
	log8.Println("Object Added:", eid, typ, x, y, z)
}

func HandleMobSpawn(c net.Conn) {
	eid, _ := ReadInt(c)
	typ, _ := ReadByte(c)
	x, _ := ReadInt(c)
	y, _ := ReadInt(c)
	z, _ := ReadInt(c)
	yaw, _ := ReadByte(c)
	pitch, _ := ReadByte(c)
	data, _ := ReadDataStream(c)
	log8.Println("Mob spawn:", eid, typ, x, y, z, yaw, pitch, data)
}

func HandlePainting(c net.Conn) {
	eid, _ := ReadInt(c)
	title, _ := ReadString(c)
	x, _ := ReadInt(c)
	y, _ := ReadInt(c)
	z, _ := ReadInt(c)
	typ, _ := ReadInt(c)
	log8.Println("Painting:", eid, title, x, y, z, typ)
}

func HandleEntityVelocity(c net.Conn) {
	eid, _ := ReadInt(c)
	vx, _ := ReadShort(c)
	vy, _ := ReadShort(c)
	vz, _ := ReadShort(c)
	log9.Println("Entity Velocity:", eid, vx, vy, vz)
}

func HandleDestroyEntity(c net.Conn) {
	eid, _ := ReadInt(c)
	log8.Println("Destory Entity:", eid)
}

func HandleEntity(c net.Conn) {
	eid, _ := ReadInt(c)
	log8.Println("Entity:", eid)
}

func HandleEntityRelativeMove(c net.Conn) {
	eid, _ := ReadInt(c)
	dx, _ := ReadByte(c)
	dy, _ := ReadByte(c)
	dz, _ := ReadByte(c)
	log9.Println("Entity Relative Move:", eid, dx, dy, dz)
}

func HandleEntityLook(c net.Conn) {
	eid, _ := ReadInt(c)
	yaw, _ := ReadByte(c)
	pitch, _ := ReadByte(c)
	log8.Println("Entity Look:", eid, yaw, pitch)
}

func HandleEntityLookAndRelativeMove(c net.Conn) {
	eid, _ := ReadInt(c)
	dx, _ := ReadByte(c)
	dy, _ := ReadByte(c)
	dz, _ := ReadByte(c)
	yaw, _ := ReadByte(c)
	pitch, _ := ReadByte(c)
	log9.Println("Entity Look And Relative Move:", eid, dx, dy, dz, yaw, pitch)
}

func HandleEntityTeleport(c net.Conn) {
	eid, _ := ReadInt(c)
	x, _ := ReadInt(c)
	y, _ := ReadInt(c)
	z, _ := ReadInt(c)
	yaw, _ := ReadByte(c)
	pitch, _ := ReadByte(c)
	log5.Println("Entity teleport:", eid, x, y, z, yaw, pitch)
}

func HandleEntityStatus(c net.Conn) {
	eid, _ := ReadInt(c)
	estatus, _ := ReadByte(c)
	log5.Println("Entity status:", eid, estatus)
}

func HandleEntityMetadata(c net.Conn) {
	eid, _ := ReadInt(c)
	data, _ := ReadDataStream(c)
	log5.Println("Entity metadata:", eid, data)
}

func HandlePreChunk(c net.Conn) {
	x, _ := ReadInt(c)
	z, _ := ReadInt(c)
	mode, _ := ReadByte(c)
	log7.Println("Pre-Chunk:", x, z, mode)
}

func HandleChunk(c net.Conn) {
	x, _ := ReadInt(c)
	y, _ := ReadShort(c)
	z, _ := ReadInt(c)
	sx, _ := ReadByte(c)
	sy, _ := ReadByte(c)
	sz, _ := ReadByte(c)
	csize, _ := ReadInt(c)
	log4.Println("Chunk received:", x, y, z, sx, sy, sz, csize)
	cdataSlice := make([]byte, csize)
	for i := int32(0); i < csize; i++ {
		b, _ := ReadByte(c)
		cdataSlice[i] = b
	}
	log7.Println("First two bytes (extranious):", cdataSlice[:2])
	cdata := bytes.NewBuffer(cdataSlice[2:])
	data := flate.NewReader(cdata)
	
	usize := int( 
		(float32(sx)+1)*
		(float32(sy)+1)*
		(float32(sz)+1))
	log7.Println("Chunck is size:", csize, "(compressed), and", float32(usize)*float32(2.5), "(uncompressed).", usize, "blocks total")
	
	blocks := make([]Block, usize)
	for i := 0; i < usize; i++ {
		blocks[i].Type, _ = ReadByte(data)
	}
	log8.Println("Read block types")
	for i := 0; i < usize; i+=2 {
		metadata, _ := ReadByte(data)
		blocks[i].Metadata = metadata >> 4
		blocks[i+1].Metadata = metadata & 0xF
	}
	log8.Println("Read block data")
	for i := 0; i < usize; i+=2 {
		light, _ := ReadByte(data)
		blocks[i].Light = light >> 4
		blocks[i+1].Light = light & 0xF
	}
	log8.Println("Read block light")
	for i := 0; i < usize; i+=2 {
		skylight, _ := ReadByte(data)
		blocks[i].Skylight = skylight >> 4
		blocks[i+1].Skylight = skylight & 0xF
	}
	UpdateChunk(&WorldPoint{x, y, z}, &Dimensions{sx, sy, sz}, blocks)
	log9.Println("Chunk blocks:", blocks)
}
/*
type MultiBlockChangeData struct {
	coords int16
	typ byte
	data byte
}*/

func HandleMultiBlockChange(c net.Conn) {
	x, _ := ReadInt(c)
	z, _ := ReadInt(c)
	size, _ := ReadShort(c)
	points := make([]LocalChunkCoords, size)
	newtypes := make([]byte, size)
	newdatas := make([]byte, size)
	for i := int16(0); i < size; i++ {
		xzcoords, _ := ReadByte(c)
		x := xzcoords >> 4
		z := xzcoords & 0xF
		y, _ := ReadByte(c)
		points[i] = LocalChunkCoords{x, y, z}
// 		arrays[i].coords, _ = ReadShort(c)
	}
	for i := int16(0); i < size; i++ {
		newtypes[i], _ = ReadByte(c)
// 		arrays[i].typ, _ = ReadByte(c)
	}
	for i := int16(0); i < size; i++ {
		newdatas[i], _ = ReadByte(c)
// 		arrays[i].data, _ = ReadByte(c)
	}
	UpdateMultipleBlocks(&ChunkCoords{x, z}, points, newtypes, newdatas)
	log7.Println("Multi-block change:", x, z, size, points, newtypes, newdatas)
}

func HandleBlockChange(c net.Conn) {
	x, _ := ReadInt(c)
	y, _ := ReadByte(c)
	z, _ := ReadInt(c)
	typ, _ := ReadByte(c)
	data, _ := ReadByte(c)
	UpdateBlock(&WorldPoint{x, int16(y), z}, typ, data)
	log7.Println("Block Change:", x, y, z, typ, data)
}

func HandlePlayNoteBlock(c net.Conn) {
	x, _ := ReadInt(c)
	y, _ := ReadShort(c)
	z, _ := ReadInt(c)
	typ, _ := ReadByte(c)
	pitch, _ := ReadByte(c)
	log5.Println("Play Note Block:", x, y, z, typ, pitch)
}

func HandleSetSlot(c net.Conn) {
	wid, _ := ReadByte(c)
	slot, _ := ReadShort(c)
	iid, _ := ReadShort(c)
	var count byte
	var uses int16
	if iid != -1 {
		count, _ = ReadByte(c)
		uses, _ = ReadShort(c)
	}
	log6.Println("Set Slot:", wid, slot, iid, count, uses)
}

func HandleWindowItems(c net.Conn) {
	wid, _ := ReadByte(c)
	count, _ := ReadShort(c)
	for i := int16(0); i < count; i++ {
		iid, _ := ReadShort(c)
		if iid != -1 {
			num2, _ := ReadByte(c)
			uses, _ := ReadShort(c)
			log6.Println("Item Update:", iid, num2, uses)
		}
		log6.Println("Empty slot:", i, iid)
	}
	log6.Println("Window Items:", wid, count)
}

func HandleSignUpdate(c net.Conn) {
	x, _ := ReadInt(c)
	y, _ := ReadShort(c)
	z, _ := ReadInt(c)
	text1, _ := ReadString(c)
	text2, _ := ReadString(c)
	text3, _ := ReadString(c)
	text4, _ := ReadString(c)
	log4.Println("Sign update:", x, y, z, text1, text2, text3, text4)
}

func HandleKick(c net.Conn) {
	reason, _ := ReadString(c)
	log0.Println("Kicked from server:", reason)
	os.Exit(0)
}