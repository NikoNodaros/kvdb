package db

import (
	"context"
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

func (kv *KVStore) Get(ctx context.Context, key string) (string, bool) {
	kv.RLock()
	defer kv.RUnlock()
	value, exists := kv.store[key]
	return value, exists
}

func (kv *KVStore) Set(ctx context.Context, key, value string) {
	kv.Lock()
	defer kv.Unlock()
	kv.store[key] = value
}

func (kv *KVStore) Delete(ctx context.Context, key string) bool {
	kv.Lock()
	defer kv.Unlock()
	if _, exists := kv.store[key]; exists {
		delete(kv.store, key)
		return true
	}
	return false
}

func (kv *KVStore) ListKeys(ctx context.Context) []string {
	kv.RLock()
	defer kv.RUnlock()
	keys := make([]string, 0, len(kv.store))
	for key := range kv.store {
		keys = append(keys, key)
	}
	return keys
}
