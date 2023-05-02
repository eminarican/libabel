package session

import "github.com/df-mc/dragonfly/server/block/cube"

type Data struct {
	MusicVol   float64
	ShowPos    bool
	ShowOthers bool

	Markers map[string]Marker
}

type Marker struct {
	Room cube.Pos
	Hex  string
	Desc string
}

func NewData() *Data {
	return &Data{
		MusicVol:   1.0,
		ShowPos:    true,
		ShowOthers: true,
		Markers:    map[string]Marker{},
	}
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
