package session

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"image"
	"time"
	_ "unsafe"
)

type Gadget struct{}

func init() {
	world.RegisterItem(Gadget{})
}

func AddGadget(p *player.Player) {
	_ = p.Inventory().SetItem(8, item.NewStack(Gadget{}, 1).
		WithCustomName(text.Colourf("<purple>Open Menu</purple>")))
}

func (s Gadget) Use(_ *world.World, user item.User, _ *item.UseContext) bool {
	p := user.(*player.Player)
	p.SendForm(newMenuForm())
	return true
}

func (Gadget) Cooldown() time.Duration {
	return time.Second / 2
}

func (Gadget) Name() string {
	return "Menu"
}

func (Gadget) Category() category.Category {
	return category.Equipment()
}

func (Gadget) EncodeItem() (name string, meta int16) {
	return "minecraft:compass", 1
}

func (Gadget) Texture() image.Image {
	return image.NewRGBA(image.Rect(0, 0, 1, 1))
}

// noinspection ALL
//
//go:linkname newMenuForm github.com/eminarican/libabel/library/menu.New
func newMenuForm() form.Form
