package session

import "github.com/df-mc/dragonfly/server/block/cube"

type State struct {
	Data *Data
	Hex  string
	Room cube.Pos
}

type Data struct {
	Markers  *Markers
	Settings *Settings
}

type Markers struct {
}

type Settings struct {
}
