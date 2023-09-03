package linearsearch

func Linearsearchv1(array []int, search int) int {
	var index = -1

	for eachIndex := 0; eachIndex < len(array); eachIndex++ {
		if array[eachIndex] == search {
			index = eachIndex
		}
	}

	return index
}

func Linearsearchv2(array []int, search int) int {
	for eachIndex := 0; eachIndex < len(array); eachIndex++ {
		if array[eachIndex] == search {
			return eachIndex
		}
	}

	return -1
}
