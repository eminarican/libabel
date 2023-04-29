package library

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/world"
)

func DefaultConfig(gen world.Generator) (server.Config, error) {
	cfg := server.DefaultConfig()

	cfg.Server.Name = "Library of Babel"
	cfg.Server.JoinMessage = ""
	cfg.Server.QuitMessage = ""
	cfg.Server.ShutdownMessage = "Library Closed"
	cfg.Server.AuthEnabled = false
	cfg.Players.SaveData = false
	cfg.World.SaveData = false
	cfg.Resources.AutoBuildPack = false

	srvCfg, err := cfg.Config(NopLogger{})

	srvCfg.Generator = func(dim world.Dimension) world.Generator {
		return gen
	}

	return srvCfg, err
}

type NopLogger struct{}

func (n NopLogger) Errorf(string, ...any) {}
func (n NopLogger) Debugf(string, ...any) {}
func (n NopLogger) Infof(string, ...any)  {}
func (n NopLogger) Fatalf(string, ...any) {}
func (n NopLogger) Warnf(string, ...any)  {}
