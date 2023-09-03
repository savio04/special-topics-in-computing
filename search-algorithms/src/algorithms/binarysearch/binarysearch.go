package binarysearch

func Binarysearch(array []int, search int, startindex int, endindex int) int {
	middle := (startindex + endindex) / 2

	if array[middle] == search {
		return middle
	}

	if startindex >= endindex {
		return -1
	}

	if array[middle] > search {
		return Binarysearch(array, search, startindex, middle-1)
	} else {
		return Binarysearch(array, search, middle+1, endindex)
	}
}
