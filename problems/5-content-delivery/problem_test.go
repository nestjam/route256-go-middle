package main

import (
	"reflect"
	"testing"
)

func Test_distribute(t *testing.T) {
	type args struct {
		weights  []int
		throuput []int
	}
	tests := []struct {
		name          string
		args          args
		wantDelta     int
		wantServers   []int
		ignoreServers bool
	}{
		{
			name: "3.5",
			args: args{
				weights:  []int{653808640, 85279389, 795940956, 795940940, 653808639},
				throuput: []int{28, 23, 28, 3, 13},
			},
			wantDelta:     0,
			wantServers:   []int{2, 4, 1, 1, 2}, //[]int{2, 4, 3, 3, 2},
			ignoreServers: true,
		},
		{
			name: "3.4",
			args: args{
				weights:  []int{98515231, 15156192, 15156192, 102304275, 15156192},
				throuput: []int{26, 4, 4, 27, 9},
			},
			wantDelta:     0,
			wantServers:   []int{1, 2, 2, 4, 2}, //[]int{1, 2, 2, 4, 3},
			ignoreServers: true,
		},
		{
			name: "3.3",
			args: args{
				weights:  []int{672441548, 593185062, 738215513, 179361325, 673832450},
				throuput: []int{11, 24, 16, 26, 5},
			},
			wantDelta:   9055669,
			wantServers: []int{2, 3, 2, 5, 2},
		},
		{
			name: "3.2",
			args: args{
				weights:  []int{958950712, 451227163, 918240076, 81129127, 21706246},
				throuput: []int{5, 14, 17, 19, 23},
			},
			wantDelta:   37352260,
			wantServers: []int{5, 2, 5, 1, 1}, //[]int{5, 2, 5, 1, 1},
		},
		{
			name: "3.1",
			args: args{
				weights:  []int{653007359, 493384397, 322725582, 648904718, 48733063},
				throuput: []int{25, 9, 9, 18, 6},
			},
			wantDelta:   17998117,
			wantServers: []int{1, 1, 4, 1, 5},
		},
		{
			name: "2.5",
			args: args{
				weights:  []int{300211984, 522442831, 88716537, 63463818, 323504321},
				throuput: []int{5, 13, 13, 19, 19},
			},
			wantDelta:   14804228,
			wantServers: []int{2, 4, 1, 1, 2},
		},
		{
			name: "2.4",
			args: args{
				weights:  []int{213611471, 403488326, 213611464, 593365178, 593365200},
				throuput: []int{17, 15, 25, 9, 11},
			},
			wantDelta:   0,
			wantServers: []int{4, 1, 4, 3, 3},
		},
		{
			name: "2.3",
			args: args{
				weights:  []int{691955117, 968737181, 968737173, 691955124, 92260684},
				throuput: []int{21, 2, 13, 15, 8},
			},
			wantDelta:   0,
			wantServers: []int{4, 1, 1, 4, 2},
		},
		{
			name: "2.2",
			args: args{
				weights:  []int{1393814, 604643178, 94597329, 935834228, 779688738},
				throuput: []int{18, 11, 24, 7, 27},
			},
			wantDelta:   34461410,
			wantServers: []int{4, 5, 5, 5, 5}, //[]int{4, 1, 4, 5, 3},
		},
		{
			name: "2.1",
			args: args{
				weights:  []int{568405993, 284202993, 603931375, 639456742, 284202993},
				throuput: []int{17, 16, 8, 24, 18},
			},
			wantDelta:   0,
			wantServers: []int{2, 3, 1, 5, 3},
		},
		{
			args: args{
				weights:  []int{12, 14, 7, 9},
				throuput: []int{3, 5},
			},
			wantDelta:   0,
			wantServers: []int{2, 2, 1, 1},
		},
		{
			args: args{
				weights:  []int{12, 13, 14, 15, 16},
				throuput: []int{3, 5},
			},
			wantDelta:   1,
			wantServers: []int{1, 1, 1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := distribute(tt.args.weights, tt.args.throuput)
			if got != tt.wantDelta {
				t.Errorf("distribute() got = %v, want %v", got, tt.wantDelta)
			}

			if !reflect.DeepEqual(got1, tt.wantServers) {
				t.Errorf("distribute() got1 = %v, want %v", got1, tt.wantServers)
			}
		})
	}
}
