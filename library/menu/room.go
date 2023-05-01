package menu

import (
	"github.com/df-mc/dragonfly/server/item"
	df "github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/algo"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
	form "github.com/twistedasylummc/inline-forms"
)

func NewRoom(ses *session.Session) df.Form {
	var btn []form.Button

	for skr := 1; skr <= 10; skr++ {
		skr := skr
		btn = append(btn, form.Button{
			Text:  text.Colourf("<dark-grey>%v.Shulker</dark-grey>", skr),
			Image: "textures/items/shulker_shell",
			Submit: func() {
				ses.SendForm(NewShulker(ses, skr))
			},
		})
	}

	return &form.Menu{
		Title:   "Room Content",
		Buttons: btn,
	}
}

func NewShulker(ses *session.Session, skr int) df.Form {
	var btn []form.Button

	for vol := 1; vol <= 10; vol++ {
		vol := vol
		btn = append(btn, form.Button{
			Text:  text.Colourf("<dark-grey>%v.Book</dark-grey>", vol),
			Image: "textures/items/book_written",
			Submit: func() {
				var pages []string

				for i := 1; i <= 50; i++ {
					pages = append(pages, algo.Read(algo.Address{
						Hex:     ses.Hex(),
						Room:    ses.Room(),
						Shulker: skr,
						Volume:  vol,
						Page:    i,
					}))
				}

				_, _ = ses.Inventory().AddItem(item.NewStack(item.WrittenBook{
					Title:  text.Colourf("<white>%v [%v] %v.Book</white>", ses.Room(), skr, vol),
					Author: "Library of Babel",
					Pages:  pages,
				}, 1))
			},
		})
	}

	return &form.Menu{
		Title:   "Shulker Content",
		Buttons: btn,
	}
}
