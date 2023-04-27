package menu

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Room struct {
	hex     string
	room    cube.Pos
	buttons []form.Button
}

func NewRoom(hex string, rom cube.Pos) form.Form {
	var btn []form.Button

	for i := 1; i <= 10; i++ {
		btn = append(btn, form.Button{
			Text:  text.Colourf("<dark-grey>%v.Shulker</dark-grey>", i),
			Image: "textures/items/shulker_shell",
		})
	}

	return form.NewMenu(Room{
		hex:     hex,
		room:    rom,
		buttons: btn,
	}, "Room Content").WithButtons(btn...)
}

func (r Room) Submit(sub form.Submitter, pressed form.Button) {
	for i, btn := range r.buttons {
		if btn != pressed {
			continue
		}

		p := sub.(*player.Player)
		p.SendForm(NewShulker(r.hex, r.room, i+1))
	}
}
