package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		want      bool
	}{
		{
			name:      "example 1 - can place 1 flower",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			name:      "example 2 - cannot place 2 flowers",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
		{
			name:      "single empty spot - can place 1",
			flowerbed: []int{0},
			n:         1,
			want:      true,
		},
		{
			name:      "all empty - can place multiple",
			flowerbed: []int{0, 0, 0, 0, 0},
			n:         3,
			want:      true,
		},
		{
			name:      "no flowers needed",
			flowerbed: []int{1, 0, 1, 0, 1},
			n:         0,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying the original slice
			flowerbed := make([]int, len(tt.flowerbed))
			copy(flowerbed, tt.flowerbed)

			got := canPlaceFlowers(flowerbed, tt.n)
			if got != tt.want {
				t.Errorf("canPlaceFlowers(%v, %d) = %v, want %v", tt.flowerbed, tt.n, got, tt.want)
			}
		})
	}
}
