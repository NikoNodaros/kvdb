package db

import (
	"context"
	"sync"
	"testing"
)

func TestKeyValueStore(t *testing.T) {
	ctx := context.Background()
	tests := map[string]func(t *testing.T){
		"TestSetAndGet": func(t *testing.T) {
			kv := New()
			kv.Set(ctx, "k", "v")
			value, exists := kv.Get(ctx, "k")
			if !exists || value != "v" {
				t.Errorf("Expected 'v', got '%s'", value)
			}
		},
		"TestDelete": func(t *testing.T) {
			kv := New()
			kv.Set(ctx, "k", "v")
			deleted := kv.Delete(ctx, "k")
			if !deleted {
				t.Errorf("Expected key 'k' to be deleted")
			}
			_, exists := kv.Get(ctx, "k")
			if exists {
				t.Errorf("Expected key 'k' to not exist")
			}
		},
		"TestListKeys": func(t *testing.T) {
			kv := New()
			kv.Set(ctx, "k1", "v1")
			kv.Set(ctx, "k2", "v2")
			kv.Set(ctx, "k3", "v3")
			keys := kv.ListKeys(ctx)
			if len(keys) != 3 {
				t.Errorf("Expected 3 keys, got %d", len(keys))
			}
		},

		"TestParallelSet": func(t *testing.T) {
			kv := New()
			var wg sync.WaitGroup
			values := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

			for _, val := range values {
				wg.Add(1)
				go func(value string) {
					defer wg.Done()
					kv.Set(ctx, "key", value)
				}(val)
			}
			wg.Wait()
			value, _ := kv.Get(ctx, "key")
			found := false
			for _, v := range values {
				if v == value {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Final value '%s' not in expected values", value)
			}
		},
	}

	for name, testFunc := range tests {
		t.Run(name, func(t *testing.T) {
			testFunc(t)
		})
	}
}
