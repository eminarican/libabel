package menu

import (
	df "github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/session"
	form "github.com/twistedasylummc/inline-forms"
)

func NewMarker(ses *session.Session) df.Form {
	btn := []form.Button{
		{
			Text: "Add New",
			Submit: func() {
				//ses.SendForm()
			},
		},
	}

	// todo for markers add btn

	return &form.Menu{
		Title:   "Marker Menu",
		Buttons: btn,
		Submit: func(closed bool) {
			if closed {
				ses.SendFormF(New)
			}
		},
	}
}
