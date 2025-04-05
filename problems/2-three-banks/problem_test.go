package threebanks

import "testing"

func Test_exchange(t *testing.T) {
	type args struct {
		b1 bank
		b2 bank
		b3 bank
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "5.5",
			args: args{
				b1: bank{
					rubleDollar: newCourse(46, 10),
					rubleEuro:   newCourse(10, 59),
					dollarRuble: newCourse(22, 89),
					dollarEuro:  newCourse(23, 47),
					euroRuble:   newCourse(7, 31),
					euroDollar:  newCourse(14, 69),
				},
				b2: bank{
					rubleDollar: newCourse(1, 92),
					rubleEuro:   newCourse(63, 56),
					dollarRuble: newCourse(11, 60),
					dollarEuro:  newCourse(25, 38),
					euroRuble:   newCourse(49, 84),
					euroDollar:  newCourse(96, 42),
				},
				b3: bank{
					rubleDollar: newCourse(3, 51),
					rubleEuro:   newCourse(92, 37),
					dollarRuble: newCourse(75, 21),
					dollarEuro:  newCourse(97, 22),
					euroRuble:   newCourse(49, 100),
					euroDollar:  newCourse(469, 85),
				},
			},
			want: 6327.090909090909,
		},
		{
			name: "1",
			args: args{
				b1: bank{
					rubleDollar: newCourse(100, 1),
					rubleEuro:   newCourse(100, 1),
					dollarRuble: newCourse(1, 100),
					dollarEuro:  newCourse(3, 2),
					euroRuble:   newCourse(1, 100),
					euroDollar:  newCourse(2, 3),
				},
				b2: bank{
					rubleDollar: newCourse(100, 1),
					rubleEuro:   newCourse(100, 1),
					dollarRuble: newCourse(1, 100),
					dollarEuro:  newCourse(3, 2),
					euroRuble:   newCourse(1, 100),
					euroDollar:  newCourse(2, 3),
				},
				b3: bank{
					rubleDollar: newCourse(100, 1),
					rubleEuro:   newCourse(100, 1),
					dollarRuble: newCourse(1, 100),
					dollarEuro:  newCourse(3, 2),
					euroRuble:   newCourse(1, 100),
					euroDollar:  newCourse(2, 3),
				},
			},
			want: 0.015,
		},
		{
			name: "3",
			args: args{
				b1: bank{
					rubleDollar: newCourse(100, 1),
					rubleEuro:   newCourse(100, 1),
					dollarRuble: newCourse(1, 100),
					dollarEuro:  newCourse(1, 1),
					euroRuble:   newCourse(1, 100),
					euroDollar:  newCourse(1, 1),
				},
				b2: bank{
					rubleDollar: newCourse(100, 1),
					rubleEuro:   newCourse(100, 1),
					dollarRuble: newCourse(1, 100),
					dollarEuro:  newCourse(1, 1),
					euroRuble:   newCourse(1, 100),
					euroDollar:  newCourse(1, 1),
				},
				b3: bank{
					rubleDollar: newCourse(100, 1),
					rubleEuro:   newCourse(100, 1),
					dollarRuble: newCourse(1, 100),
					dollarEuro:  newCourse(1, 1),
					euroRuble:   newCourse(1, 100),
					euroDollar:  newCourse(1, 1),
				},
			},
			want: 0.01,
		},
		{
			name: "4",
			args: args{
				b1: bank{
					rubleDollar: newCourse(1, 10),
					rubleEuro:   newCourse(1, 100),
					dollarRuble: newCourse(1, 100),
					dollarEuro:  newCourse(1, 100),
					euroRuble:   newCourse(1, 100),
					euroDollar:  newCourse(1, 100),
				},
				b2: bank{
					rubleDollar: newCourse(1, 100),
					rubleEuro:   newCourse(1, 100),
					dollarRuble: newCourse(1, 100),
					dollarEuro:  newCourse(1, 100),
					euroRuble:   newCourse(1, 100),
					euroDollar:  newCourse(1, 100),
				},
				b3: bank{
					rubleDollar: newCourse(1, 100),
					rubleEuro:   newCourse(1, 100),
					dollarRuble: newCourse(1, 100),
					dollarEuro:  newCourse(1, 100),
					euroRuble:   newCourse(1, 100),
					euroDollar:  newCourse(1, 100),
				},
			},
			want: 1e+06,
		},
		{
			name: "2",
			args: args{
				b1: bank{
					rubleDollar: newCourse(42, 68),
					rubleEuro:   newCourse(35, 1),
					dollarRuble: newCourse(70, 25),
					dollarEuro:  newCourse(79, 59),
					euroRuble:   newCourse(63, 65),
					euroDollar:  newCourse(6, 46),
				},
				b2: bank{
					rubleDollar: newCourse(82, 28),
					rubleEuro:   newCourse(62, 92),
					dollarRuble: newCourse(96, 43),
					dollarEuro:  newCourse(28, 37),
					euroRuble:   newCourse(92, 5),
					euroDollar:  newCourse(3, 54),
				},
				b3: bank{
					rubleDollar: newCourse(93, 83),
					rubleEuro:   newCourse(22, 17),
					dollarRuble: newCourse(19, 96),
					dollarEuro:  newCourse(48, 27),
					euroRuble:   newCourse(72, 39),
					euroDollar:  newCourse(70, 13),
				},
			},
			want: 16.392857142857142,
		},
		{
			name: "5.2",
			args: args{
				b1: bank{
					rubleDollar: newCourse(68, 100),
					rubleEuro:   newCourse(36, 95),
					dollarRuble: newCourse(4, 12),
					dollarEuro:  newCourse(23, 34),
					euroRuble:   newCourse(74, 65),
					euroDollar:  newCourse(42, 12),
				},
				b2: bank{
					rubleDollar: newCourse(54, 69),
					rubleEuro:   newCourse(48, 45),
					dollarRuble: newCourse(63, 58),
					dollarEuro:  newCourse(38, 60),
					euroRuble:   newCourse(24, 42),
					euroDollar:  newCourse(30, 79),
				},
				b3: bank{
					rubleDollar: newCourse(17, 36),
					rubleEuro:   newCourse(91, 43),
					dollarRuble: newCourse(89, 7),
					dollarEuro:  newCourse(41, 43),
					euroRuble:   newCourse(65, 49),
					euroDollar:  newCourse(47, 6),
				},
			},
			want: 9.779411764705882,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchange([...]*bank{&tt.args.b1, &tt.args.b2, &tt.args.b3}); got != tt.want {
				t.Errorf("exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}
