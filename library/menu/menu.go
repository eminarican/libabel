package menu

import (
	"fmt"
	df "github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
	form "github.com/twistedasylummc/inline-forms"
	"golang.design/x/clipboard"
)

func New(ses *session.Session) df.Form {
	btn := []form.Button{
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
				ses.SendFormF(NewMarker)
			},
		},
		{
			Text:  "Clear\nclears inventory",
			Image: "textures/items/flint_and_steel",
			Submit: func() {
				ses.SendFormF(NewClear)
			},
		},
	}

	if ses.Local() {
		btn = append(btn, form.Button{
			Text:  "Copy\ncopies session data",
			Image: "textures/items/banner_pattern",
			Submit: func() {
				rom := ses.Room()

				clipboard.Write(clipboard.FmtText, []byte(fmt.Sprintf(
					"%v %v %v %v", rom.X(), rom.Y(), rom.Z(), ses.Hex(),
				)))

				ses.Message(text.Colourf("<green>Session data copied to clipboard</green>"))
			},
		})
	}

	return &form.Menu{
		Title:   "Menu",
		Content: "Please select an option below",
		Buttons: btn,
	}
}
