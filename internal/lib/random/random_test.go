package random

import (
	"testing"
)

func TestRandomURL(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size = 1",
			size: 1,
		},
		{
			name: "size = 5",
			size: 5,
		},
		{
			name: "size = 10",
			size: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := RandomURL(tt.size)
			if len(result) != tt.size {
				t.Errorf("RandomURL() returned wrong length: got %d, want %d", len(result), tt.size)
			}
		})
	}
}
