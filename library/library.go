package library

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/eminarican/libabel/library/command"
	_ "github.com/eminarican/libabel/library/menu"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Library struct {
	s *server.Server
}

func New(c server.Config) *Library {
	return &Library{
		s: c.New(),
	}
}

func (l *Library) Accept(p *player.Player) {
	p.Message(text.Colourf("<yellow>Welcome to Library of Babel!</yellow>"))
	p.SetGameMode(session.GameMode{})

	session.AddGadget(p)

	p.Handle(&session.Session{
		Server: l.s,
		Player: p,
		State: &session.State{
			Hex: "libraryofbabel",
		},
	})
	p.Inventory().Handle(&session.InvHandler{Player: p})
}

func (l *Library) Start() {
	l.s.CloseOnProgramEnd()
	l.s.Listen()

	l.setupWorld()

	cmd.Register(cmd.New(
		"tp", "Teleports you to specified place or player",
		nil, command.TeleportRoom{}, command.TeleportPlayer{},
	))

	fmt.Println(text.ANSI(text.Colourf(
		"<green>Library of Babel Started on Port 19132!</green>",
	)))

	for l.s.Accept(l.Accept) {
	}
}

func (l *Library) setupWorld() {
	w := l.s.World()

	w.StopTime()
	w.StopRaining()
	w.StopThundering()
	w.StopWeatherCycle()

	w.SetTime(18000)
	w.SetSpawn(cube.Pos{8, 7*16 + 1, 8})
}
