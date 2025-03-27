package tests

import (
	"leetcode-go/solutions"
	"reflect"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		number int
		want   string
	}{
		{3749, "MMMDCCXLIX"},
	}

	for _, tt := range tests {
		got := solutions.IntToRoman(tt.number)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("IntToRoman(%v) = %v", tt.number, tt.want)
		}
	}
}
