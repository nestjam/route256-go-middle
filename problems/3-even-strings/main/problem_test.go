package main

import "testing"

func Test_findPairs(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				[]string{
					"xxxiy",
					"xxzj",
					"xxyix",
				},
			},
			want: 1,
		},
		{
			args: args{
				[]string{
					"ab",
					"cb",
					"cd",
					"cb",
				},
			},
			want: 5,
		},
		{
			name: "8",
			args: args{
				[]string{
					"zxzxzxzxzxzxzxzxzxzt",
					"zxzxzxzxzxzxzxzxzxzk",
				},
			},
			want: 1,
		},
		{
			name: "3.1",
			args: args{
				[]string{
					"yxcxwxnx",
				},
			},
			want: 0,
		},
		{
			name: "2.1",
			args: args{
				[]string{
					"gxpx",
					"bx",
					"yxexk",
					"y",
					"axex",
				},
			},
			want: 3,
		},
		{
			name: "1.6",
			args: args{
				[]string{
					"a",
					"a",
				},
			},
			want: 1,
		},
		{
			name: "1.7",
			args: args{
				[]string{
					"a",
					"b",
				},
			},
			want: 0,
		},
		{
			name: "1.4",
			args: args{
				[]string{
					"aaaa",
					"aaaa",
					"aaaa",
					"aaa",
				},
			},
			want: 6,
		},
		{
			name: "1.3",
			args: args{
				[]string{
					"abca",
					"abc",
				},
			},
			want: 1,
		},
		{
			name: "1.2",
			args: args{
				[]string{
					"asd",
					"das",
					"sda",
				},
			},
			want: 0,
		},
		{
			name: "1.1",
			args: args{
				[]string{
					"ababa",
					"ababa",
					"ababa",
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairs(tt.args.strings); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
