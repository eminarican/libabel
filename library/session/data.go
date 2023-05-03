package session

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Data struct {
	Markers map[string]Marker `json:"markers"`
}

type Marker struct {
	Room cube.Pos `json:"room"`
	Hex  string   `json:"hex"`
	Desc string   `json:"desc"`
}

func (m Marker) Format() string {
	return text.Colourf(
		"<green>Room:</green> %v\n"+
			"<purple>Desc:</purple> %v\n"+
			"<yellow>Hex:</yellow> %v",
		m.Room, m.Desc, m.Hex,
	)
}

func NewData() *Data {
	return &Data{
		Markers: map[string]Marker{},
	}
}

func (s *Session) Data() *Data {
	return s.data
}

func (s *Session) SetData(dat *Data) {
	s.data = dat
}

func (s *Session) Markers() map[string]Marker {
	return s.data.Markers
}

func (s *Session) Marker(name string) Marker {
	return s.data.Markers[name]
}

func (s *Session) SetMarker(name string, mar Marker) {
	s.data.Markers[name] = mar
}

func (s *Session) RemMarker(name string) {
	delete(s.data.Markers, name)
}
