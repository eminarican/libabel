package session

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Data struct {
	MusicVol   float64 `json:"music_vol"`
	ShowPos    bool    `json:"show_pos"`
	ShowOthers bool    `json:"show_others"`

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
		MusicVol:   1.0,
		ShowPos:    true,
		ShowOthers: true,
		Markers:    map[string]Marker{},
	}
}

func (s *Session) Data() *Data {
	return s.data
}

func (s *Session) SetData(dat *Data) {
	s.data = dat
}

func (s *Session) MusicVol() float64 {
	return s.data.MusicVol
}

func (s *Session) SetMusicVol(new float64) {
	s.data.MusicVol = new
}

func (s *Session) ShowPos() bool {
	return s.data.ShowPos
}

func (s *Session) SetShowPos(new bool) {
	s.data.ShowPos = new
}

func (s *Session) ShowOthers() bool {
	return s.data.ShowOthers
}

func (s *Session) SetShowOthers(new bool) {
	s.data.ShowOthers = new
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
