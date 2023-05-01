package menu

import (
	df "github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
	form "github.com/twistedasylummc/inline-forms"
)

func New(ses *session.Session) df.Form {
	return &form.Menu{
		Title:   "Menu",
		Content: "Please select an item",
		Buttons: []form.Button{
			{
				Text:  "Search\nopens search menu",
				Image: "textures/items/ender_eye",
				Submit: func() {
					ses.SendFormF(NewSearch)
				},
			},
			{
				Text:  "Marker\nopens marker menu",
				Image: "textures/items/book_portfolio",
				Submit: func() {
					ses.Message(text.Colourf("<red>This is not yet implemented</red>"))
					//ses.SendFormF(NewMarker)
				},
			},
			{
				Text:  "Clear\nclears inventory",
				Image: "textures/items/flint_and_steel",
				Submit: func() {
					ses.SendFormF(NewClear)
				},
			},
			{
				Text:  "Settings\nopens settings menu",
				Image: "textures/items/banner_pattern",
				Submit: func() {
					ses.Message(text.Colourf("<red>This is not yet implemented</red>"))
					//ses.SendFormF(NewSettings)
				},
			},
		},
	}
}
