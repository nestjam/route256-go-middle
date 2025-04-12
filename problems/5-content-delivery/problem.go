package main

func distribute(imageWeights []int, serverThrouput []int) (minDelta int, imageStorages []int) {
	imageCount := len(imageWeights)
	times := make([][]int, imageCount)

	for i := 0; i < imageCount; i++ {
		times[i] = make([]int, len(serverThrouput))
		for j := 0; j < len(serverThrouput); j++ {
			t := imageWeights[i] / serverThrouput[j]

			if imageWeights[i]%serverThrouput[j] > 0 {
				t++
			}

			times[i][j] = t
		}
	}

	imageStorages = make([]int, imageCount)
	combination := make([]int, imageCount)
	minDelta = getDelta(combination, times)
	serverCount := len(serverThrouput)

	for {
		combination[imageCount-1]++

		for i := imageCount - 1; i > 0; i-- {
			if combination[i]/serverCount > 0 {
				combination[i] = 0
				combination[i-1]++
			} else {
				break
			}
		}

		if combination[0]/serverCount > 0 {
			break
		}

		delta := getDelta(combination, times)

		if minDelta > delta {
			minDelta = delta
			copy(imageStorages, combination)
		}

		if minDelta == 0 {
			break
		}
	}

	for i := 0; i < imageCount; i++ {
		imageStorages[i]++
	}

	return
}

func getDelta(servers []int, times [][]int) int {
	server := servers[0]
	max, min := times[0][server], times[0][server]

	for i := 1; i < len(times); i++ {
		server = servers[i]
		if max < times[i][server] {
			max = times[i][server]
		}

		if min > times[i][server] {
			min = times[i][server]
		}
	}

	return max - min
}
