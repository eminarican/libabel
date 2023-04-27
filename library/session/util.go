package session

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

func roomPosFromVec3(vec3 mgl64.Vec3) cube.Pos {
	return cube.Pos{
		int(math.Floor(vec3[0])) >> 4,
		int(vec3[1] / 16),
		int(math.Floor(vec3[2])) >> 4,
	}
}
