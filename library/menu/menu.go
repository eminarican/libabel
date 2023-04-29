package menu

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Menu struct {
	Search   form.Button
	Marker   form.Button
	Clear    form.Button
	Settings form.Button
}

func New() form.Form {
	return form.NewMenu(Menu{
		Search:   form.NewButton("Search\nopens search menu", "textures/items/ender_eye"),
		Marker:   form.NewButton("Marker\nopens marker menu", "textures/items/book_portfolio"),
		Clear:    form.NewButton("Clear\nclears inventory", "textures/items/flint_and_steel"),
		Settings: form.NewButton("Settings\nopens settings menu", "textures/items/banner_pattern"),
	}, "Menu")
}

func (m Menu) Submit(sub form.Submitter, pressed form.Button) {
	p := sub.(*player.Player)

	switch pressed {
	case m.Marker:
		p.Message(text.Colourf("<red>Not yet implemented!</red>"))
	case m.Search:
		p.SendForm(NewSearch())
	case m.Clear:
		p.SendForm(NewClear())
	case m.Settings:
		p.Message(text.Colourf("<red>Not yet implemented!</red>"))
	}
}
