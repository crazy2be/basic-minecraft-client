package main

type Block struct {
	Type byte
	Metadata byte
	Light byte
	Skylight byte
}

type ChunkCoords struct {
	X int32
	Z int32
}

// Coordinates of a block within a chunk. X and Z are actually only nibbles or half-bytes.
type LocalChunkCoords struct {
	X byte
	Y byte
	Z byte
}

type WorldPoint struct {
	X int32
	Y int16
	Z int32
}

type Dimensions struct {
	X byte
	Y byte
	Z byte
}

// Indexed by chunk x, chunk z, then local (within the chunk) x, y, and z location.
type chunkData [16][128][16]Block
var worldData = make(map[int32]map[int32]*chunkData)

// Updates a bunch of blocks within a chunk
func UpdateChunk(p *WorldPoint, d *Dimensions, data []Block) {
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
				//log.Println("index:", index, x, y, z)
				chunk[x][y][z] = data[index]
				index ++
			}
		}
	}
}

func UpdateBlock(p *WorldPoint, newtype, metadata byte) {
	cc := p.ChunkCoords()
	lcc := p.LocalChunkCoords()
	chunk := AllocChunk(cc)
	block := &chunk[lcc.X][lcc.Y][lcc.Z]
	block.Type = newtype
	block.Metadata = metadata
}

func UpdateMultipleBlocks(cc *ChunkCoords, points []LocalChunkCoords, newtypes, newdatas []byte) {
	for i := 0; i < len(points); i++ {
		p := cc.WorldCoords(&points[i])
		UpdateBlock(p, newtypes[i], newdatas[i])
	}
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

// Allocates a chunk if one does not exist at the specified ChunkCoords.
func AllocChunk(cc *ChunkCoords) *chunkData {
	if !ChunkExists(cc) {
		worldData[cc.X] = make(map[int32]*chunkData)
		worldData[cc.X][cc.Z] = new(chunkData)
	}
	return worldData[cc.X][cc.Z]
}

// Gets the coordinates of the chunk that the point (in world coords) lives within
func (p *WorldPoint) ChunkCoords() *ChunkCoords {
	// Is Y always 0?
	return &ChunkCoords{p.X >> 4, p.Z >> 4}
}

// Gets the local coordinates (with the chunk) where the point (in world coords) is.
func (p *WorldPoint) LocalChunkCoords() *LocalChunkCoords {
	return &LocalChunkCoords{byte(p.X & 15), byte(p.Y & 127), byte(p.Z & 15)}
}

func (cc *ChunkCoords) WorldCoords(lcc *LocalChunkCoords) *WorldPoint {
	wx := cc.X << 4 | int32(lcc.X)
	wy := int16(lcc.Y)
	wz := cc.Z << 4 | int32(lcc.Z)
	return &WorldPoint{wx, wy, wz}
}