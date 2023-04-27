package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Clear struct{}

func (Clear) Run(src cmd.Source, out *cmd.Output) {
	p := src.(*player.Player)

	p.Inventory().Clear()

	out.Print(text.Colourf("<green>Inventory Cleared</green>"))
}
