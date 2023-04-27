package menu

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/eminarican/libabel/library/page"
	"github.com/eminarican/libabel/library/session"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Search struct {
	Text  form.Input
	Exact form.Toggle
}

func NewSearch() form.Form {
	return form.New(Search{
		Text:  form.NewInput("Please fill the form to search library", "", "text to search"),
		Exact: form.NewToggle("Exact Match", false),
	}, "Search")
}

func (c Search) Submit(sub form.Submitter) {
	p := sub.(*player.Player)
	txt := c.Text.Value()

	if len(txt) > page.Length {
		p.Message(text.Colourf("<red>Text can't be longer than page length %v character</red>", page.Length))
		return
	}

	adr := page.Search(txt, !c.Exact.Value())
	p.SendForm(NewSearchResult(adr))
}

type SearchResult struct {
	address  page.Address
	Teleport form.Button
}

func NewSearchResult(adr page.Address) form.Form {
	return form.NewMenu(SearchResult{
		address:  adr,
		Teleport: form.NewButton("Teleport", ""),
	}, "Search Result").WithBody(
		text.Colourf(" <green>Room:</green> %v\n", adr.Room[:]),
		text.Colourf("<purple>Shulker:</purple> %v\n", adr.Shulker),
		text.Colourf("<aqua>Book:</aqua> %v\n", adr.Volume),
		text.Colourf("<red>Page:</red> %v\n", adr.Page),
		text.Colourf("<yellow>Hex:</yellow> %v\n", adr.Hex),
	)
}

func (c SearchResult) Submit(sub form.Submitter, pressed form.Button) {
	if pressed != c.Teleport {
		return
	}

	p := sub.(*player.Player)
	h := p.Handler().(*session.Handler)

	h.Room = c.address.Room
	h.Hex = c.address.Hex

	p.Teleport(h.Room.Vec3().Mul(16).Add(mgl64.Vec3{8, 1, 8}))
	p.Message(text.Colourf("<green>Teleported to room</green>"))
	p.Message(text.Colourf("<purple>Shulker:</purple> %v\n", c.address.Shulker))
	p.Message(text.Colourf("<aqua>Book:</aqua> %v\n", c.address.Volume))
	p.Message(text.Colourf("<red>Page:</red> %v\n", c.address.Page))
}
