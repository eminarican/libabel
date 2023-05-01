package menu

import (
	df "github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/session"
	form "github.com/twistedasylummc/inline-forms"
)

func NewSettings(ses *session.Session) df.Form {
	music := 0.0
	pos := true

	return &form.Custom{
		Title: "Settings",
		Elements: []form.Element{
			form.Slider{
				Text:     "Music Volume",
				Min:      0,
				Max:      100,
				StepSize: 10,
				Default:  music,
				Submit: func(value float64) {
					music = value
				},
			},
			form.Toggle{
				Text:    "Show Position",
				Default: pos,
				Submit: func(enabled bool) {
					pos = enabled
				},
			},
		},
		Submit: func(closed bool, _ []any) {
			if closed {
				ses.SendFormF(New)
				return
			}
		},
	}
}
