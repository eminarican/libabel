package menu

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Clear struct {
	Yes form.Button
	No  form.Button
}

func NewClear() form.Form {
	return form.NewModal(Clear{
		Yes: form.YesButton(),
		No:  form.NoButton(),
	}, "Clear Confirmation").
		WithBody("are you sure about to clear your inventory?")
}

func (c Clear) Submit(sub form.Submitter, pressed form.Button) {
	p := sub.(*player.Player)

	switch pressed {
	case c.Yes:
		p.Inventory().Clear()
		session.AddGadget(p)
		p.Message(text.Colourf("<green>Inventory Cleared</green>"))
	case c.No:
		p.SendForm(New())
	}
}
