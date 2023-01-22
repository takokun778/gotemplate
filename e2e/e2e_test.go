//go:build e2e
// +build e2e

package internal

import (
	"net/http"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res, err := http.Get("http://localhost:8080")
			if err != nil {
				t.Errorf("http.Get() error = %v", err)
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("http.Get() status = %v", res.StatusCode)
			}
		})
	}
}
