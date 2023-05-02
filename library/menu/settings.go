package menu

import (
	df "github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/session"
	form "github.com/twistedasylummc/inline-forms"
)

func NewSettings(ses *session.Session) df.Form {
	music := ses.MusicVol()
	other := ses.ShowOthers()
	pos := ses.ShowPos()

	return &form.Custom{
		Title: "Settings",
		Elements: []form.Element{
			form.Label{
				Text: "Please update your settings",
			},
			form.Toggle{
				Text:    "Show Position",
				Default: pos,
				Submit: func(enabled bool) {
					pos = enabled
				},
			},
			form.Toggle{
				Text:    "Show Players",
				Default: other,
				Submit: func(enabled bool) {
					other = enabled
				},
			},
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
		},
		Submit: func(closed bool, _ []any) {
			if closed {
				ses.SendFormF(New)
				return
			}

			ses.SetMusicVol(music)
			ses.SetShowOthers(other)
			ses.SetShowPos(pos)

			if !ses.ShowPos() {
				ses.RemoveScoreboard()
			}
		},
	}
}
