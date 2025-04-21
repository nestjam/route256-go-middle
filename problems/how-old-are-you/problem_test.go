package main

import (
	"reflect"
	"testing"
)

func Test_youngerFact_calc(t *testing.T) {
	type args struct {
		op    string
		facts []fact
	}
	tests := []struct {
		name    string
		f       youngerFact
		args    args
		want    int
		wantErr bool
	}{
		{
			args: args{
				op: "A",
				facts: []fact{
					youngerFact{
						id:    0,
						op:    "A",
						op2:   "B",
						value: 2,
					},
					constFact{
						id:    1,
						op:    "B",
						value: 12,
					},
				},
			},
			f: youngerFact{
				id:    0,
				op:    "A",
				op2:   "B",
				value: 2,
			},
			want:    10,
			wantErr: false,
		},
		{
			args: args{
				op: "A",
				facts: []fact{
					youngerFact{
						id:    0,
						op:    "B",
						op2:   "A",
						value: 2,
					},
					constFact{
						id:    1,
						op:    "A",
						value: 10,
					},
				},
			},
			f: youngerFact{
				id:    0,
				op:    "B",
				op2:   "A",
				value: 2,
			},
			want:    10,
			wantErr: false,
		},
		{
			args: args{
				op: "A",
				facts: []fact{
					youngerFact{
						id:    0,
						op:    "B",
						op2:   "A",
						value: 2,
					},
					constFact{
						id:    1,
						op:    "B",
						value: 8,
					},
				},
			},
			f: youngerFact{
				id:    0,
				op:    "B",
				op2:   "A",
				value: 2,
			},
			want:    10,
			wantErr: false,
		},
		{
			args: args{
				op: "A",
				facts: []fact{
					youngerFact{
						id:    0,
						op:    "A",
						op2:   "B",
						value: 2,
					},
					constFact{
						id:    1,
						op:    "A",
						value: 10,
					},
				},
			},
			f: youngerFact{
				id:    0,
				op:    "A",
				op2:   "B",
				value: 2,
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.calc(tt.args.op, tt.args.facts)
			if (err != nil) != tt.wantErr {
				t.Errorf("youngerFact.calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("youngerFact.calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equalityFact_calc(t *testing.T) {
	type args struct {
		op    string
		facts []fact
	}
	tests := []struct {
		name    string
		f       equalityFact
		args    args
		want    int
		wantErr bool
	}{
		{
			args: args{
				op: "A",
				facts: []fact{
					equalityFact{
						id:  0,
						op:  "B",
						op2: "A",
					},
					constFact{
						id:    1,
						op:    "B",
						value: 10,
					},
				},
			},
			f: equalityFact{
				id:  0,
				op:  "B",
				op2: "A",
			},
			want:    10,
			wantErr: false,
		},
		{
			args: args{
				op: "A",
				facts: []fact{
					equalityFact{
						id:  0,
						op:  "A",
						op2: "B",
					},
					constFact{
						id:    1,
						op:    "A",
						value: 10,
					},
				},
			},
			f: equalityFact{
				id:  0,
				op:  "A",
				op2: "B",
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.calc(tt.args.op, tt.args.facts)
			if (err != nil) != tt.wantErr {
				t.Errorf("equalityFact.calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("equalityFact.calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseConstFact(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want constFact
	}{
		{
			args: args{
				"A is 10 years old",
			},
			want: constFact{
				op:    "A",
				value: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFact([]byte(tt.args.s)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTarget(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "How old is I?",
			},
			want: "I",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTarget(tt.args.s); got != tt.want {
				t.Errorf("parseTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
