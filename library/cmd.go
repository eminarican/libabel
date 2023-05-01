package library

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type TeleportRoom struct {
	X   int                  `cmd:"x"`
	Y   int                  `cmd:"y"`
	Z   int                  `cmd:"z"`
	Hex cmd.Optional[string] `cmd:"hex"`
}

func (t TeleportRoom) Run(src cmd.Source, out *cmd.Output) {
	ses := session.Get(src.(*player.Player))

	if t.Y < 1 || t.Y > 14 {
		out.Error("room y position should be in range of 1-14")
		return
	}

	ses.SetRoom(cube.Pos{
		t.X, t.Y, t.Z,
	})
	if hex, ok := t.Hex.Load(); ok {
		ses.SetHex(hex)
	}

	ses.TeleportRoom(ses.Room())
	out.Print(text.Colourf(
		"<green>Teleported to Room: %v Hex: %v</green>", ses.Room(), ses.Hex(),
	))
}

type TeleportPlayer struct {
	Target []cmd.Target `cmd:"target"`
}

func (t TeleportPlayer) Run(src cmd.Source, out *cmd.Output) {
	ses := session.Get(src.(*player.Player))

	if len(t.Target) > 1 {
		out.Errorf("You can't select multiple targets")
		return
	}

	tar, ok := t.Target[0].(*player.Player)
	if !ok {
		out.Errorf("Target isn't a player")
		return
	}

	if tar == ses.Player() {
		out.Errorf("You can't teleport to yourself")
		return
	}

	tarSes := session.Get(tar)
	ses.SetRoom(tarSes.Room())
	ses.SetHex(tarSes.Hex())

	ses.Teleport(tar.Position())
	out.Print(text.Colourf("<green>Teleported to Player: %v</green>", tar.Name()))
}
