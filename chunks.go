package main

import (
	"fmt"
)

type Block struct {
	Type byte
	Metadata byte
	Light byte
	Skylight byte
}

func (b *Block) String() string {
	return fmt.Sprintf("%s block (block id: %d). Additional: %d, %d, %d", b.NiceName(), b.Type, b.Metadata, b.Light, b.Skylight)
}

type ChunkDimensions struct {
	X byte
	Y byte
	Z byte
}

// Indexed by chunk x, chunk z, then local (within the chunk) x, y, and z location.
type chunkData [16][128][16]Block
var worldData = make(map[int32]map[int32]*chunkData)

// Updates a bunch of blocks within a chunk
func UpdateChunk(p *BlockCoords, d *ChunkDimensions, data []Block) {
	cc := p.ChunkCoords()
	log6.Println("data length:", len(data))
	lcc := p.LocalChunkCoords()
	// 0 0 0
	// 4 0 -2
	// 64 0 -32
	// 15 127 15
	log7.Println("localChunkCoords:", lcc, "chunkCoords:", cc, "worldPoint:", p, "dimensions:", d)
	chunk := AllocChunk(cc)
	index := 0
	for x := int32(lcc.X); x < int32(d.X)+1+int32(lcc.X); x++ {
		for z := int32(lcc.Z); z < int32(d.Z)+1+int32(lcc.Z); z++ {
			for y := int16(lcc.Y); y < int16(d.Y)+1+int16(lcc.Y); y++ {
				//index := int32(y)+(z*(int32(d.Y)+1))+(x*(int32(d.Y)+1))+(int32(d.Z)+1)
				//index := (x+1)*(z+1)*(int32(y)+1)
				if x > 15 || y > 127 || z > 15 {
					log0.Println("WARNING: MALFORMED CHUNK DATA!")
					log0.Println("IMPENDING CRASH!")
					log0.Println("index:", index, x, y, z)
					log0.Println("Mitigiating crash by ignoring chunk data...")
					return
				}
				chunk[x][y][z] = data[index]
				index ++
			}
		}
	}
}

func UpdateBlock(bc *BlockCoords, newtype, metadata byte) {
	cc := bc.ChunkCoords()
	lcc := bc.LocalChunkCoords()
	chunk := AllocChunk(cc)
	block := &chunk[lcc.X][lcc.Y][lcc.Z]
	block.Type = newtype
	block.Metadata = metadata
}

func UpdateMultipleBlocks(cc *ChunkCoords, points []LocalChunkCoords, newtypes, newdatas []byte) {
	for i := 0; i < len(points); i++ {
		p := cc.BlockCoords(&points[i])
		UpdateBlock(p, newtypes[i], newdatas[i])
	}
}

func GetBlock(bc *BlockCoords) *Block {
	cc := bc.ChunkCoords()
	lcc := bc.LocalChunkCoords()
	chunk := GetChunk(cc)
	if chunk == nil {
		return nil
	}
	block := chunk[lcc.X][lcc.Y][lcc.Z]
	return &block
}

func ChunkExists(cc *ChunkCoords) bool {
	_, ok := worldData[cc.X]
	if !ok {
		return false
	}
	_, ok = worldData[cc.X][cc.Z]
	if !ok {
		return false
	}
	return true
}

func GetChunk(cc *ChunkCoords) *chunkData {
	if !ChunkExists(cc) {
		return nil
	}
	return worldData[cc.X][cc.Z]
}

// Allocates a chunk if one does not exist at the specified ChunkCoords.
func AllocChunk(cc *ChunkCoords) *chunkData {
	if !ChunkExists(cc) {
		worldData[cc.X] = make(map[int32]*chunkData)
		worldData[cc.X][cc.Z] = new(chunkData)
	}
	return worldData[cc.X][cc.Z]
}