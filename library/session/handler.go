package session

import (
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
	_ "unsafe"
)

var (
	height = mgl64.Vec3{0, 224, 0}
)

type Handler struct {
	player.NopHandler
	Hex    string
	Room   cube.Pos
	Player *player.Player
}

func (h *Handler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, _ float64, _ float64) {
	p := h.Player

	newRom := roomPosFromVec3(newPos)
	delRom := newRom.Sub(h.Room)

	if newRom != h.Room {
		if newRom.Y() < 1 {
			p.Teleport(newPos.Add(height))
			h.Room[1] = 14
			ctx.Cancel()
		} else if newRom.Y() > 14 {
			p.Teleport(newPos.Sub(height))
			h.Room[1] = 1
			ctx.Cancel()
		} else {
			h.Room[1] += delRom.Y()
		}

		h.Room[0] += delRom.X()
		h.Room[2] += delRom.Z()
	}

	p.SendScoreboard(scoreboard.New(text.Colourf("<white>%v</white>", h.Room[:])))
}

func (h *Handler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, _ cube.Face, _ mgl64.Vec3) {
	p := h.Player

	if _, ok := p.World().Block(pos).(block.EnchantingTable); ok {
		p.SendForm(newRoomForm(h.Hex, h.Room))
	}

	ctx.Cancel()
}

func (h *Handler) HandleItemDrop(ctx *event.Context, _ *entity.Item) {
	ctx.Cancel()
}

func (h *Handler) HandleFoodLoss(ctx *event.Context, _ int, _ *int) {
	ctx.Cancel()
}

func (h *Handler) HandleItemPickup(ctx *event.Context, _ item.Stack) {
	ctx.Cancel()
}

// noinspection ALL
//
//go:linkname newRoomForm github.com/eminarican/libabel/library/menu.NewRoom
func newRoomForm(hex string, rom cube.Pos) form.Form
