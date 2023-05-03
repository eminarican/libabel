package menu

import (
	df "github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/algo"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
	form "github.com/twistedasylummc/inline-forms"
)

func NewSearch(ses *session.Session) df.Form {
	txt := ""
	xct := false

	return &form.Custom{
		Title: "Search",
		Elements: []form.Element{
			form.Label{
				Text: "Please fill the form to search library",
			},
			form.Input{
				Text:        "Text:",
				Placeholder: "example input",
				Submit: func(text string) {
					txt = text
				},
			},
			form.Toggle{
				Text:    "Exact Match",
				Default: false,
				Submit: func(enabled bool) {
					xct = enabled
				},
			},
		},
		Submit: func(closed bool, _ []any) {
			if closed {
				ses.SendFormF(New)
				return
			}

			if len(txt) > algo.Length {
				ses.Message(text.Colourf(
					"<red>Text can't be longer than page length %v character</red>", algo.Length,
				))
				return
			}

			adr := algo.Search(txt, !xct)
			ses.SendForm(NewSearchResult(ses, adr))
		},
	}
}

func NewSearchResult(ses *session.Session, adr algo.Address) df.Form {
	return &form.Menu{
		Title:   "Search Result",
		Content: adr.Format(),
		Buttons: []form.Button{
			{
				Text: "Teleport",
				Submit: func() {
					ses.SetHex(adr.Hex)
					ses.TeleportRoom(adr.Room)

					ses.Message(text.Colourf("<green>Teleported to destination</green>"))
					ses.Message(adr.Format())
				},
			},
		},
	}
}
