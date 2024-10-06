package db

import (
	"testing"
)

func TestKeyValueStore(t *testing.T) {
	tests := map[string]func(t *testing.T){
		"TestSetAndGet": func(t *testing.T) {
		},

		"TestDelete": func(t *testing.T) {
		},

		"TestListKeys": func(t *testing.T) {
		},

		"TestParallelSet": func(t *testing.T) {
		},
	}

	for name, testFunc := range tests {
		t.Run(name, func(t *testing.T) {
			testFunc(t)
		})
	}
}
