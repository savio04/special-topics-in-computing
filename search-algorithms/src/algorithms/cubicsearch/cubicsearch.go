package cubicsearch

func Cubicsearch(array []int, search int) int {
	size := len(array)
	position := -1

	for primaryindex := 0; primaryindex < size; primaryindex++ {
		for secondindex := 0; secondindex < size; secondindex++ {
			for thirdindex := 0; thirdindex < size; thirdindex++ {
				if array[primaryindex] == search && array[secondindex] == search && array[thirdindex] == search {
					position = primaryindex
				}
			}
		}
	}

	return position
}
