package main

import (
	"math"
	"slices"
)

func distribute(images []int, servers []int) (minDelta int, storages []int) {
	imageGroups := make(map[int][]int)

	for i := 0; i < len(images); i++ {
		imageGroups[images[i]] = append(imageGroups[images[i]], i)
	}

	uniqueImages := make([]int, len(imageGroups))
	i := 0
	for weight, _ := range imageGroups {
		uniqueImages[i] = weight
		i++
	}
	
	times := make([][]int, len(uniqueImages))

	for i := 0; i < len(uniqueImages); i++ {
		times[i] = make([]int, len(servers))
		for j := 0; j < len(servers); j++ {
			t := uniqueImages[i] / servers[j]

			if uniqueImages[i]%servers[j] > 0 {
				t++
			}

			times[i][j] = t
		}
	}

	minDelta = math.MaxInt

	for i := 0; i < len(uniqueImages); i++ {
		for j := 0; j < len(servers); j++ {
			ref := times[i][j]
			s, delta := find(times, ref)

			if delta == 0 {
				return delta, ungroupImages(s, uniqueImages, imageGroups, images)
			}

			if minDelta > delta {
				minDelta = delta
				storages = s
			}
		}
	}

	return minDelta, ungroupImages(storages, uniqueImages, imageGroups, images)
}

func ungroupImages(storages []int, uniqueImages []int, weightGroups map[int][]int, images []int) []int{
	newStorages := make([]int, len(images))

	for i := 0; i < len(storages); i++ {
		server := storages[i]
		imageWeight := uniqueImages[i]
		imagesGroup := weightGroups[imageWeight]
		for j := 0; j < len(imagesGroup); j++ {
			newStorages[imagesGroup[j]] = server
		}
	}

	return newStorages
}

func find(times [][]int, ref int) (storages []int, minDelta int) {
	intervals := make([]interval, len(times))

	for i := 0; i < len(times); i++ {
		interval := interval{
			image: i,
			left:  deliveryTime{server: -1, time: math.MinInt},
			right: deliveryTime{server: -1, time: math.MaxInt},
		}

		for j := 0; j < len(times[i]); j++ {
			value := times[i][j] - ref

			if value == 0 {
				interval.left.server = j
				interval.left.time = times[i][j]
				interval.right = interval.left
				break
			} else if value < 0 && times[i][j] > interval.left.time {
				interval.left.server = j
				interval.left.time = times[i][j]
			} else if value > 0 && times[i][j] < interval.right.time {
				interval.right.server = j
				interval.right.time = times[i][j]
			}
		}

		intervals[i] = interval
	}

	slices.SortFunc(intervals, func(a, b interval) int {
		return a.getAbsMin(ref) - b.getAbsMin(ref)
	})

	min, max, storages := findRecursive(0, intervals, ref, ref, ref, make([]int, len(intervals)))

	for i := 0; i < len(storages); i++ {
		storages[i]++
	}

	return storages, max - min
}

func findRecursive(i int, intervals []interval, refTime, minTime, maxTime int, storages []int) (int, int, []int) {
	if i == len(intervals) {
		return minTime, maxTime, storages
	}

	interval := intervals[i]

	if interval.left.server == -1 {
		storages[interval.image] = interval.right.server
		return findRecursive(i+1, intervals, refTime, minTime, max(maxTime, interval.right.time), storages)
	} else if interval.right.server == -1 { 
		storages[interval.image] = interval.left.server
		return findRecursive(i+1, intervals, refTime, min(minTime, interval.left.time), maxTime, storages)
	} else if abs(refTime-interval.left.time) == abs(refTime-interval.right.time) && refTime-interval.right.time != 0 {
		storagesOp1 := make([]int, len(storages))
		copy(storagesOp1, storages)
		storagesOp1[interval.image] = interval.left.server

		minOp1, maxOp1, storagesOp1 := findRecursive(i+1, intervals, refTime, min(minTime, interval.left.time), maxTime, storagesOp1)

		storagesOp2 := make([]int, len(storages))
		copy(storagesOp2, storages)
		storagesOp2[interval.image] = interval.right.server

		minOp2, maxOp2, storagesOp2 := findRecursive(i+1, intervals, refTime, minTime, max(maxTime, interval.right.time), storagesOp2)

		if maxOp1-minOp1 < maxOp2-minOp2 {
			return minOp1, maxOp1, storagesOp1
		} else {
			return minOp2, maxOp2, storagesOp2
		}
	} else if maxTime - interval.left.time < interval.right.time - minTime {
		storages[interval.image] = interval.left.server
		return findRecursive(i+1, intervals, refTime, min(minTime, interval.left.time), maxTime, storages)
	} else {
		storages[interval.image] = interval.right.server
		return findRecursive(i+1, intervals, refTime, minTime, max(maxTime, interval.right.time), storages)
	}
}

type interval struct {
	image int
	left, right deliveryTime
}

type deliveryTime struct {
	server int
	time   int
}

func (i interval) getAbsMin(ref int) int {
	return min(abs(ref-i.left.time), abs(ref-i.right.time))
}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}