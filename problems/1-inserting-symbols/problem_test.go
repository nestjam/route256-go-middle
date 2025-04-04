package insertingsymbols

import "testing"

func Test_checkString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "abab",
			},
			want: false,
		},
		{
			args: args{
				s: "PppP",
			},
			want: false,
		},
		{
			args: args{
				s: "aba",
			},
			want: true,
		},
		{
			args: args{
				s: "aba",
			},
			want: true,
		},
		{
			args: args{
				s: "baa",
			},
			want: false,
		},
		{
			args: args{
				s: "aa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkString(tt.args.s); got != tt.want {
				t.Errorf("checkString() = %v, want %v", got, tt.want)
			}
		})
	}
}
