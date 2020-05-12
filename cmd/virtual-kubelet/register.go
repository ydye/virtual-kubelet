package main

import (
	"fmt"
	"github.com/virtual-kubelet/virtual-kubelet/cmd/virtual-kubelet/internal/provider"
	"github.com/virtual-kubelet/virtual-kubelet/cmd/virtual-kubelet/internal/provider/mock"
)

func registerMock(s *provider.Store) {
	s.Register("mock", func(cfg provider.InitConfig) (provider.Provider, error) { //nolint:errcheck
		fmt.Println("[debug][point2][internal] >>>>>>>")
		return mock.NewMockProvider(
			cfg.ConfigPath,
			cfg.NodeName,
			cfg.OperatingSystem,
			cfg.InternalIP,
			cfg.DaemonPort,
		)
	})
}
