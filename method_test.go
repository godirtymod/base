package base

import "testing"

func TestApple(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Case1", "ZZZ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Apple(); got == tt.want {
				t.Errorf("Apple() = %v, unwanted %v", got, tt.want)
			}
		})
	}
}
