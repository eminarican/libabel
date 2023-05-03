package library

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	_ "github.com/eminarican/libabel/library/menu"
	"github.com/eminarican/libabel/library/session"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"log"
)

type Library struct {
	srv *server.Server
	pro *session.Provider
}

func New(c server.Config, pro *session.Provider) *Library {
	return &Library{
		srv: c.New(),
		pro: pro,
	}
}

func (l *Library) HandleJoin(ses *session.Session) {
	if dat, err := l.pro.Load(ses.UUID()); err == nil {
		ses.SetData(&dat)
	}
	ses.Message(text.Colourf("<yellow>Welcome to Library of Babel!</yellow>"))
}

func (l *Library) HandleQuit(ses *session.Session) {
	if err := l.pro.Save(ses.UUID(), *ses.Data()); err != nil {
		log.Fatalf("player data couldn't be saved: %v %v", err, ses.Data())
	}
}

func (l *Library) Start() {
	l.srv.CloseOnProgramEnd()
	l.srv.Listen()

	l.setupWorld()

	cmd.Register(cmd.New(
		"tp", "Teleports you to specified place or player",
		nil, TeleportRoom{}, TeleportPlayer{},
	))

	log.Println(text.ANSI(text.Colourf(
		"<green>Library of Babel Started on port 19132!</green>",
	)))

	for l.srv.Accept(func(p *player.Player) {
		l.HandleJoin(session.Init(p, l.srv, l.HandleQuit))
	}) {
	}
}

func (l *Library) setupWorld() {
	w := l.srv.World()

	w.StopTime()
	w.StopRaining()
	w.StopThundering()
	w.StopWeatherCycle()

	w.SetTime(18000)
	w.SetSpawn(cube.Pos{8, 7*16 + 1, 8})
}
