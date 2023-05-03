package menu

import (
	df "github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
	form "github.com/twistedasylummc/inline-forms"
)

func NewClear(ses *session.Session) df.Form {
	return &form.Modal{
		Title:   "Clear Confirmation",
		Content: "are you sure about to clear your inventory?",
		Button1: form.Button{
			Text: "Yes",
			Submit: func() {
				inv := ses.Inventory()

				inv.Clear()
				session.AddGadget(inv)

				ses.Message(text.Colourf("<green>Inventory cleared</green>"))
			},
		},
		Button2: form.Button{
			Text: "No",
			Submit: func() {
				ses.SendFormF(New)
			},
		},
	}
}
