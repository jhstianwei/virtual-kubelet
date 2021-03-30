package main

import (
	"github.com/jhstianwei/virtual-kubelet/providers"
	"github.com/jhstianwei/virtual-kubelet/providers/mock"
)

func registerMock(s *providers.Store) {
	s.Register("mock", func(cfg providers.InitConfig) (providers.Provider, error) {
		return mock.NewMockProvider(
			cfg.ConfigPath,
			cfg.NodeName,
			cfg.OperatingSystem,
			cfg.InternalIP,
			cfg.DaemonPort,
		)
	})

	s.Register("mockV0", func(cfg providers.InitConfig) (providers.Provider, error) {
		return mock.NewMockProvider(
			cfg.ConfigPath,
			cfg.NodeName,
			cfg.OperatingSystem,
			cfg.InternalIP,
			cfg.DaemonPort,
		)
	})
}
