package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{1}, want: 1},
		{args: args{2}, want: 2},
		{args: args{3}, want: 3},
		{args: args{6}, want: 13},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
