package db

import (
	"sync"
)

type KVStore struct {
	sync.RWMutex
	store map[string]string
}

func New() *KVStore {
	return &KVStore{
		store: make(map[string]string),
	}
}
