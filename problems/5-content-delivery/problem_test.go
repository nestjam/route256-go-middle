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
		name        string
		args        args
		wantDelta   int
		wantServers []int
	}{
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
