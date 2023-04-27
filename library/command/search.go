package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/eminarican/libabel/library/menu"
)

type Search struct{}

func (Search) Run(src cmd.Source, _ *cmd.Output) {
	p := src.(*player.Player)
	p.SendForm(menu.NewSearch())
}
