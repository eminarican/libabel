package library

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
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

func (l *Library) HandleJoin(ses *session.Session) {
	ses.Message(text.Colourf("<yellow>Welcome to Library of Babel!</yellow>"))
}

func (l *Library) HandleQuit(ses *session.Session) {
}

func (l *Library) Start() {
	l.s.CloseOnProgramEnd()
	l.s.Listen()

	l.setupWorld()

	cmd.Register(cmd.New(
		"tp", "Teleports you to specified place or player",
		nil, TeleportRoom{}, TeleportPlayer{},
	))

	fmt.Println(text.ANSI(text.Colourf(
		"<green>Library of Babel Started on Port 19132!</green>",
	)))

	for l.s.Accept(func(p *player.Player) {
		l.HandleJoin(session.Init(p, l.s, l.HandleQuit))
	}) {
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
