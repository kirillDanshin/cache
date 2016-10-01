package cache

import "sync"

type Cache struct {
	storage map[string]interface{}
	lock    sync.RWMutex
}
