package main

import "testing"

func TestSumString(t *testing.T) {

	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "sum all",
			args: []string{"1", "2", "3"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumString(tt.args); got != tt.want {
				t.Errorf("SumString() = %v, want %v", got, tt.want)
			}
		})
	}
}
