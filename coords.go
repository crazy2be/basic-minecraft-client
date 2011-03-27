package main

// Used for a few rare cases, such as spawn position. Represents the location of an item in absolute pixels. Less precise than WorldCoords, but more precise than BlockCoords.
type WorldPxCoords struct {
	X int32
	Y int32
	Z int32
}

func (wpc *WorldPxCoords) WorldCoords() *WorldCoords {
	return &WorldCoords{float64(wpc.X)/32, float64(wpc.Y)/32, float64(wpc.Z)/32}
}

// Coordinates of a object, in "block coordinates". Used for most things that aren't WorldCoords.
type BlockCoords struct {
	X int32
	Y byte
	Z int32
}

func (bc *BlockCoords) WorldPxCoords() *WorldPxCoords {
	return &WorldPxCoords{bc.X*32, int32(bc.Y)*32, bc.Z*32}
}

func (bc *BlockCoords) WorldCoords() *WorldCoords {
	return &WorldCoords{float64(bc.X), float64(bc.Y), float64(bc.Z)}
}

// Gets the coordinates of the chunk that the point (in world coords) lives within
func (bc *BlockCoords) ChunkCoords() *ChunkCoords {
	// Is Y always 0?
	return &ChunkCoords{bc.X >> 4, bc.Z >> 4}
}

// Gets the local coordinates (with the chunk) where the point (in world coords) is.
func (bc *BlockCoords) LocalChunkCoords() *LocalChunkCoords {
	return &LocalChunkCoords{byte(bc.X & 15), byte(bc.Y & 127), byte(bc.Z & 15)}
}


type WorldCoords struct {
	X float64
	Y float64
	Z float64
}

func (wc *WorldCoords) WorldPxCoords() *WorldPxCoords {
	return &WorldPxCoords{int32(wc.X*32), int32(wc.Y*32), int32(wc.Z*32)}
}

func (wc *WorldCoords) BlockCoords() *BlockCoords {
	return &BlockCoords{int32(wc.X), uint8(wc.Y), int32(wc.Z)}
}

// Represents the coordinates of a chunk, in world space. No Y value because chunks always extend from floor to ceiling.
type ChunkCoords struct {
	X int32
	Z int32
}

func (cc *ChunkCoords) BlockCoords(lcc *LocalChunkCoords) *BlockCoords {
	wx := cc.X << 4 | int32(lcc.X)
	wy := uint8(lcc.Y)
	wz := cc.Z << 4 | int32(lcc.Z)
	return &BlockCoords{wx, wy, wz}
}


// Coordinates of a block within a chunk. X and Z are actually only nibbles or half-bytes.
type LocalChunkCoords struct {
	X byte
	Y byte
	Z byte
}