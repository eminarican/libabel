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

func (i InvHandler) HandleTake(ctx *event.Context, _ int, stk item.Stack) {
	if _, ok := stk.Item().(Gadget); ok {
		ctx.Cancel()
	}
}

func (i InvHandler) HandleDrop(ctx *event.Context, _ int, stk item.Stack) {
	if _, ok := stk.Item().(Gadget); !ok {
		_ = i.Player.Inventory().RemoveItem(stk)
	}

	ctx.Cancel()
}
