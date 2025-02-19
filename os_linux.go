//go:build linux
// +build linux

package wgctrl

import (
	"github.com/tahirmahm123/awgctrl/internal/wginternal"
	"github.com/tahirmahm123/awgctrl/internal/wglinux"
	"github.com/tahirmahm123/awgctrl/internal/wguser"
)

// newClients configures wginternal.Clients for Linux systems.
func newClients() ([]wginternal.Client, error) {
	var clients []wginternal.Client

	// Linux has an in-kernel WireGuard implementation. Determine if it is
	// available and make use of it if so.
	kc, ok, err := wglinux.New()
	if err != nil {
		return nil, err
	}
	if ok {
		clients = append(clients, kc)
	}

	// Although it isn't recommended to use userspace implementations on Linux,
	// it can be used. We make use of it in integration tests as well.
	uc, err := wguser.New()
	if err != nil {
		return nil, err
	}

	// Kernel devices seem to appear first in wg(8).
	clients = append(clients, uc)
	return clients, nil
}
