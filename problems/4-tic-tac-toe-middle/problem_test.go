package main

import "testing"

func Test_canCrossWin(t *testing.T) {
	type args struct {
		k     int
		board [][]cell
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				k:     2,
				board: [][]cell{{x, o, e}, {o, e, e}, {e, e, e}},
			},
			want: false,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{o, x, e}, {e, o, e}, {e, e, o}},
			},
			want: false,
		},
		{
			args: args{
				k:     3,
				board: [][]cell{{o, e, x}, {o, e, e}, {o, e, x}},
			},
			want: false,
		},
		{
			args: args{
				k:     3,
				board: [][]cell{{o, o, o}, {e, e, e}, {x, x, e}},
			},
			want: false,
		},
		{
			args: args{
				k:     3,
				board: [][]cell{{o, o, x}, {o, e, o}, {x, e, e}},
			},
			want: true,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{o, x, o}, {x, o, o}, {e, e, e}},
			},
			want: false,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{e, x}, {x, e}},
			},
			want: false,
		},
		{
			args: args{
				k:     3,
				board: [][]cell{{x, o, e}, {o, e, o}, {e, e, x}},
			},
			want: true,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{o, x, o}, {o, o, x}, {e, e, e}},
			},
			want: false,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{x, e}, {e, x}},
			},
			want: false,
		},
		{
			args: args{
				k:     3,
				board: [][]cell{{x, e, o}, {x, e, x}, {o, o, e}},
			},
			want: true,
		},
		{
			args: args{
				k:     3,
				board: [][]cell{{x}, {x}, {e}},
			},
			want: true,
		},
		{
			args: args{
				k:     3,
				board: [][]cell{{x}, {e}, {x}},
			},
			want: true,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{e}, {o}, {x}},
			},
			want: false,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{e}, {x}},
			},
			want: true,
		},
		{
			args: args{
				k:     1,
				board: [][]cell{{x}, {e}},
			},
			want: false,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{x}, {e}},
			},
			want: true,
		},
		{
			args: args{
				k:     3,
				board: [][]cell{{x, x, e}},
			},
			want: true,
		},
		{
			args: args{
				k:     3,
				board: [][]cell{{x, e, x}},
			},
			want: true,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{e, o, x}},
			},
			want: false,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{e, x}},
			},
			want: true,
		},
		{
			args: args{
				k:     1,
				board: [][]cell{{x, e}},
			},
			want: false,
		},
		{
			args: args{
				k:     2,
				board: [][]cell{{x, e}},
			},
			want: true,
		},
		{
			args: args{
				k:     1,
				board: [][]cell{{o}},
			},
			want: false,
		},
		{
			args: args{
				k:     1,
				board: [][]cell{{x}},
			},
			want: false,
		},
		{
			args: args{
				k:     1,
				board: [][]cell{{e}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := newBoard(tt.args.board)
			if got := canCrossWin(&b, tt.args.k); got != tt.want {
				t.Errorf("canCrossWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
