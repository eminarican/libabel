package library

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
	"github.com/df-mc/structure"
)

type Generator struct {
	data [16][16][16]uint32
}

func NewGenerator(str structure.Structure) Generator {
	gen := Generator{}

	gen.LoopChunk(func(x uint8, y uint8, z uint8) {
		b, _ := str.At(int(x), int(y), int(z), nil)
		gen.data[x][y][z] = world.BlockRuntimeID(b)
	})

	return gen
}

func (g Generator) GenerateChunk(_ world.ChunkPos, chunk *chunk.Chunk) {
	go func() {
		g.LoopChunk(func(x uint8, y uint8, z uint8) {
			blk := g.data[x][y][z]

			for i := uint8(0); i < 16; i++ {
				chunk.SetBlock(x, int16(i*16+y), z, 0, blk)
			}
		})
	}()
}

func (g Generator) LoopChunk(callback func(x uint8, y uint8, z uint8)) {
	for x := uint8(0); x < 16; x++ {
		for y := uint8(0); y < 16; y++ {
			for z := uint8(0); z < 16; z++ {
				callback(x, y, z)
			}
		}
	}
}
