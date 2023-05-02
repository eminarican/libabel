package menu

import (
	df "github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
	form "github.com/twistedasylummc/inline-forms"
)

func NewMarker(ses *session.Session) df.Form {
	btn := []form.Button{
		{
			Text: "Add New",
			Submit: func() {
				ses.SendFormF(NewMarkerAdd)
			},
		},
	}

	for name, mar := range ses.Markers() {
		name, mar := name, mar
		btn = append(btn, form.Button{
			Text:  name,
			Image: "textures/items/paper",
			Submit: func() {
				ses.SendForm(NewMarkerView(ses, name, mar))
			},
		})
	}

	return &form.Menu{
		Title:   "Marker Menu",
		Content: "Please select an option below",
		Buttons: btn,
		Submit: func(closed bool) {
			if closed {
				ses.SendFormF(New)
			}
		},
	}
}

func NewMarkerAdd(ses *session.Session) df.Form {
	var name string
	var desc string

	return &form.Custom{
		Title: "Add Marker",
		Elements: []form.Element{
			form.Label{
				Text: "Please fill the form to add marker",
			},
			form.Input{
				Text:        "Name:",
				Placeholder: "my marker",
				Submit: func(text string) {
					name = text
				},
			},
			form.Input{
				Text:        "Description:",
				Placeholder: "look at shulker 5 book 3 page 2",
				Submit: func(text string) {
					desc = text
				},
			},
			form.Label{
				Text: text.Colourf("<green>Room:</green> %v\n<yellow>Hex</yellow> %v", ses.Room(), ses.Hex()),
			},
		},
		Submit: func(closed bool, _ []any) {
			defer ses.SendFormF(NewMarker)

			if closed {
				return
			}

			ses.SetMarker(name, session.Marker{
				Room: ses.Room(),
				Hex:  ses.Hex(),
				Desc: desc,
			})
		},
	}
}

func NewMarkerView(ses *session.Session, name string, mar session.Marker) df.Form {
	return &form.Menu{
		Title:   "View Marker",
		Content: text.Colourf("<red>Name:</red> %v\n%v", name, mar.Format()),
		Buttons: []form.Button{
			{
				Text:  "Teleport",
				Image: "textures/items/ender_pearl",
				Submit: func() {
					ses.SetHex(mar.Hex)
					ses.TeleportRoom(mar.Room)

					ses.Message(text.Colourf("<green>Teleported to Marker %v</green>", name))
				},
			},
			{
				Text:  "Delete",
				Image: "textures/items/flint_and_steel",
				Submit: func() {
					ses.SendForm(NewMarkerDelete(ses, name, mar))
				},
			},
		},
		Submit: func(closed bool) {
			if closed {
				ses.SendFormF(NewMarker)
			}
		},
	}
}

func NewMarkerDelete(ses *session.Session, name string, mar session.Marker) df.Form {
	return &form.Modal{
		Title:   "Delete Marker",
		Content: text.Colourf("<red>Name:</red> %v\n%v", name, mar.Format()),
		Button1: form.Button{
			Text: "Yes",
			Submit: func() {
				ses.RemMarker(name)
				ses.Message(text.Colourf("<green>Marker Removed</green>"))
			},
		},
		Button2: form.Button{
			Text: "No",
			Submit: func() {
				ses.SendForm(NewMarkerView(ses, name, mar))
			},
		},
	}
}
