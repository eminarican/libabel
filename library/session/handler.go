package session

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
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

func (s *Session) HandleMove(ctx *event.Context, newPos mgl64.Vec3, _ float64, _ float64) {
	newRom := roomPosFromVec3(newPos)
	delRom := newRom.Sub(s.Room())

	if newRom != s.Room() {
		if newRom.Y() < 1 {
			s.Teleport(newPos.Add(height))
			s.room[1] = 14
			ctx.Cancel()
		} else if newRom.Y() > 14 {
			s.Teleport(newPos.Sub(height))
			s.room[1] = 1
			ctx.Cancel()
		} else {
			s.room[1] += delRom.Y()
		}

		s.room[0] += delRom.X()
		s.room[2] += delRom.Z()
	}

	if s.ShowPos() {
		s.SendScoreboard(scoreboard.New(
			text.Colourf("<white>%v</white>", s.Room()),
		))
	}
}

func (s *Session) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, _ cube.Face, _ mgl64.Vec3) {
	if _, ok := s.World().Block(pos).(block.EnchantingTable); ok {
		s.SendFormF(newRoomForm)
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

func (s *Session) HandleQuit() {
	s.quit(s)
}

type InvHandler struct {
	inventory.NopHandler
	player *player.Player
}

func (i InvHandler) HandleTake(ctx *event.Context, _ int, stk item.Stack) {
	if _, ok := stk.Item().(Gadget); ok {
		ctx.Cancel()
	}
}

func (i InvHandler) HandleDrop(ctx *event.Context, _ int, stk item.Stack) {
	if _, ok := stk.Item().(Gadget); !ok {
		_ = i.player.Inventory().RemoveItem(stk)
	}

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
func newRoomForm(ses *Session) form.Form
