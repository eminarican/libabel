package command

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/eminarican/libabel/library/session"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type TeleportRoom struct {
	Room mgl64.Vec3           `cmd:"room"`
	Hex  cmd.Optional[string] `cmd:"hex"`
}

func (t TeleportRoom) Run(src cmd.Source, out *cmd.Output) {
	p := src.(*player.Player)
	h := p.Handler().(*session.Handler)

	if t.Room.Y() < 1 || t.Room.Y() > 14 {
		out.Error("room y position should be in range of 1-14")
		return
	}

	h.Room = cube.PosFromVec3(t.Room)
	if hex, ok := t.Hex.Load(); ok {
		h.Hex = hex
	}

	p.Teleport(h.Room.Vec3().Mul(16).Add(mgl64.Vec3{8, 1, 8}))
	out.Print(text.Colourf("<green>Teleported to Room: %v Hex: %v</green>", h.Room[:], h.Hex))
}

type TeleportPlayer struct {
	Target []cmd.Target `cmd:"target"`
}

func (t TeleportPlayer) Run(src cmd.Source, out *cmd.Output) {
	p := src.(*player.Player)
	h := p.Handler().(*session.Handler)

	if len(t.Target) > 1 {
		out.Errorf("You can't select multiple targets")
		return
	}

	tp, ok := t.Target[0].(*player.Player)
	if !ok {
		out.Errorf("Target isn't a player")
		return
	}

	if p == tp {
		out.Errorf("You can't teleport to yourself")
		return
	}

	th := tp.Handler().(*session.Handler)
	h.Room = th.Room
	h.Hex = th.Hex

	p.Teleport(tp.Position())
	out.Print(text.Colourf("<green>Teleported to Player: %v</green>", tp.Name()))
}