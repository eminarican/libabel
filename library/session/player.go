package session

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

func (s *Session) Inventory() *inventory.Inventory {
	return s.Player().Inventory()
}

func (s *Session) Message(a ...any) {
	s.Player().Message(a...)
}

func (s *Session) Messagef(f string, a ...any) {
	s.Player().Messagef(f, a...)
}

func (s *Session) SendForm(f form.Form) {
	s.Player().SendForm(f)
}

func (s *Session) SendFormF(f func(ses *Session) form.Form) {
	s.Player().SendForm(f(s))
}

func (s *Session) SendScoreboard(scb *scoreboard.Scoreboard) {
	s.Player().SendScoreboard(scb)
}

func (s *Session) RemoveScoreboard() {
	s.Player().RemoveScoreboard()
}

func (s *Session) Teleport(pos mgl64.Vec3) {
	s.Player().Teleport(pos)
}

func (s *Session) TeleportRoom(pos cube.Pos) {
	s.Player().Teleport(pos.Vec3().Mul(16).Add(mgl64.Vec3{8.5, 1, 8.5}))
}

func (s *Session) World() *world.World {
	return s.Player().World()
}
