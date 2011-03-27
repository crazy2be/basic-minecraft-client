package main

import (
	"net"
)

// Player-related functions, including inventory status, position, health, etc.

type PlayerCoords struct {
	// Normal coordinate structure within the world, used for all entities.
	WorldCoords
	// YTop is the top of the player's bounding box, sometimes refered to as "stance"
	YTop float64
	OnGround bool
}

type PlayerLook struct {
	Yaw float32
	Pitch float32
	OnGround bool
}

var playerPosition *PlayerCoords
var playerLook *PlayerLook

func UpdatePlayerPos(pp *PlayerCoords) {
	if pp.YTop - pp.Y < 0.1 {
		log3.Println("Illegal stance!")
		pp.YTop = pp.Y + 0.2
	}
	if pp.YTop - pp.Y > 1.65 {
		log3.Println("Illegal stance!")
		pp.YTop = pp.Y + 1.5
	}
	playerPosition = pp
}

func UpdatePlayerLook(pl *PlayerLook) {
	playerLook = pl
}

func SetPlayerPos(c net.Conn, pp *PlayerCoords) {
	UpdatePlayerPos(pp)
	SendPlayerPos(c, pp)
}

func GetPlayerPos() *PlayerCoords {
	tmp := *playerPosition
	return &tmp
}