package session

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
)

type InvHandler struct {
	inventory.NopHandler
	Player *player.Player
}

func (i InvHandler) HandleDrop(ctx *event.Context, _ int, it item.Stack) {
	_ = i.Player.Inventory().RemoveItem(it)
	ctx.Cancel()
}
