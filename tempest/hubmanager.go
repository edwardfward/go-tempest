package tempest

import "sync"

// HubManager manages Tempest hubs on a network.
type HubManager struct {
	Hubs map[string]Hub
	mu   sync.RWMutex
}
