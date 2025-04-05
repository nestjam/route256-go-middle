package threebanks

func exchange(banks [3]*bank) float32 {
	const oneRuble int32 = 1
	var maxAmount float32

	for i := 0; i < len(banks); i++ {
		amount := banks[i].rubleDollar.exchange(oneRuble)
		if amount > maxAmount {
			maxAmount = amount
		}

		for j := 0; j < len(banks); j++ {
			if i == j {
				continue
			}

			amount := banks[i].rubleEuro.combine(banks[j].euroDollar).exchange(oneRuble)
			if amount > maxAmount {
				maxAmount = amount
			}

			rubleEuroRuble := banks[i].rubleEuro.combine(banks[j].euroRuble)
			rubleDollarRuble := banks[i].rubleDollar.combine(banks[j].dollarRuble)
			rubleDollarEuro := banks[i].rubleDollar.combine(banks[j].dollarEuro)
			
			for k := 0; k < len(banks); k++ {
				if j == k || i == k {
					continue
				}
	
				courses := [...]course {
					rubleEuroRuble.combine(banks[k].rubleDollar),
					rubleDollarRuble.combine(banks[k].rubleDollar),
					rubleDollarEuro.combine(banks[k].euroDollar),
				}

				for l := 0; l < len(courses); l++ {
					amount := courses[l].exchange(oneRuble)
					if amount > maxAmount {
						maxAmount = amount
					}
				}
			}
		}
	}

	return maxAmount
}

type bank struct {
	rubleDollar course
	rubleEuro course
	dollarRuble course
	dollarEuro course
	euroRuble course
	euroDollar course
}

type course struct {
	from int32
	to int32
}

func newCourse(from int32, to int32) course {
	return course{from, to}
}

func (c course) combine(b  course) course {
	return newCourse(c.from * b.from, c.to * b.to)
}

func (c course) exchange(amount int32) float32 {
	return float32(amount * c.to) / float32(c.from)
}