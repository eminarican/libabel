package session

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
)

type Session struct {
	player.NopHandler
	quit func(session *Session)

	player *player.Player
	server *server.Server

	room cube.Pos
	hex  string
}

func Get(p *player.Player) *Session {
	return p.Handler().(*Session)
}

func Init(p *player.Player, s *server.Server, quit func(session *Session)) *Session {
	ses := &Session{
		quit:   quit,
		player: p,
		server: s,
		hex:    "library_of_babel",
	}

	p.Handle(ses)
	p.SetGameMode(GameMode{})

	p.Inventory().Handle(&InvHandler{player: p})
	AddGadget(p.Inventory())

	return ses
}

func (s *Session) Player() *player.Player {
	return s.player
}

func (s *Session) Server() *server.Server {
	return s.server
}

func (s *Session) Room() cube.Pos {
	return s.room
}

func (s *Session) SetRoom(new cube.Pos) {
	s.room = new
}

func (s *Session) Hex() string {
	return s.hex
}

func (s *Session) SetHex(new string) {
	s.hex = new
}
