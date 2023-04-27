package menu

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/page"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Shulker struct {
	hex     string
	room    cube.Pos
	shulker int
	buttons []form.Button
}

func NewShulker(hex string, rom cube.Pos, skr int) form.Form {
	var btn []form.Button

	for i := 1; i <= 10; i++ {
		btn = append(btn, form.Button{
			Text:  text.Colourf("<dark-grey>%v.Book</dark-grey>", i),
			Image: "textures/items/book_written",
		})
	}

	return form.NewMenu(Shulker{
		hex:     hex,
		room:    rom,
		shulker: skr,
		buttons: btn,
	}, "Shulker Content").WithButtons(btn...)
}

func (s Shulker) Submit(sub form.Submitter, pressed form.Button) {
	for vol, btn := range s.buttons {
		if btn != pressed {
			continue
		}

		p := sub.(*player.Player)

		var pages []string

		for i := 1; i <= 50; i++ {
			pages = append(pages, page.Read(page.Address{
				Hex:     s.hex,
				Room:    s.room,
				Shulker: s.shulker,
				Volume:  vol + 1,
				Page:    i,
			}))
		}

		_, _ = p.Inventory().AddItem(item.NewStack(item.WrittenBook{
			Title:  text.Colourf("<white>%v [%v] %v.Book</white>", s.room[:], s.shulker, vol),
			Author: "Library of Babel",
			Pages:  pages,
		}, 1))
	}
}
