package session

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"math"
	_ "unsafe"
)

var (
	height = mgl64.Vec3{0, 224, 0}
)

type Session struct {
	player.NopHandler
	State  *State
	Player *player.Player
	Server *server.Server
}

func (s *Session) HandleMove(ctx *event.Context, newPos mgl64.Vec3, _ float64, _ float64) {
	p := s.Player
	sta := s.State

	newRom := roomPosFromVec3(newPos)
	delRom := newRom.Sub(sta.Room)

	if newRom != sta.Room {
		if newRom.Y() < 1 {
			p.Teleport(newPos.Add(height))
			sta.Room[1] = 14
			ctx.Cancel()
		} else if newRom.Y() > 14 {
			p.Teleport(newPos.Sub(height))
			sta.Room[1] = 1
			ctx.Cancel()
		} else {
			sta.Room[1] += delRom.Y()
		}

		sta.Room[0] += delRom.X()
		sta.Room[2] += delRom.Z()
	}

	p.SendScoreboard(scoreboard.New(text.Colourf("<white>%v</white>", sta.Room[:])))
}

func (s *Session) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, _ cube.Face, _ mgl64.Vec3) {
	p := s.Player
	sta := s.State

	if _, ok := p.World().Block(pos).(block.EnchantingTable); ok {
		p.SendForm(newRoomForm(sta.Hex, sta.Room))
	}

	ctx.Cancel()
}

func (s *Session) HandleItemDrop(ctx *event.Context, _ *entity.Item) {
	ctx.Cancel()
}

func (s *Session) HandleFoodLoss(ctx *event.Context, _ int, _ *int) {
	ctx.Cancel()
}

func (s *Session) HandleItemPickup(ctx *event.Context, _ item.Stack) {
	ctx.Cancel()
}

func roomPosFromVec3(vec3 mgl64.Vec3) cube.Pos {
	return cube.Pos{
		int(math.Floor(vec3[0])) >> 4,
		int(vec3[1] / 16),
		int(math.Floor(vec3[2])) >> 4,
	}
}

// noinspection ALL
//
//go:linkname newRoomForm github.com/eminarican/libabel/library/menu.NewRoom
func newRoomForm(hex string, rom cube.Pos) form.Form
