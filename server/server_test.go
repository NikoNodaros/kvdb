package server

import (
	"testing"
)

func TestServer(t *testing.T) {
	tests := map[string]func(t *testing.T){
		"TestGet": func(t *testing.T) {
		},

		"TestPut": func(t *testing.T) {
		},

		"TestDelete": func(t *testing.T) {
		},

		"TestListKeys": func(t *testing.T) {
		},

		"TestParallelPuts": func(t *testing.T) {
		},
	}

	for name, testFunc := range tests {
		t.Run(name, func(t *testing.T) {
			testFunc(t)
		})
	}
}
