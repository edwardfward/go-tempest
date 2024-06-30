package tempest

import "sync"

// HubManager manages Tempest hubs on a network.
type HubManager struct {
	Hubs map[string]Hub // Map of hub serial numbers to hubs.
	mu   sync.RWMutex   // Mutex to protect the map.
}
