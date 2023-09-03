package quadraticsearch

func Quadraticsearch(array []int, search int) int {
	position := -1
	camein := false
	count := 0

	for primaryindex := 0; primaryindex < len(array); primaryindex++ {
		for secondindex := primaryindex; secondindex < len(array); secondindex++ {
			if array[primaryindex] == search {
				if !camein {
					position = primaryindex

					if array[secondindex] == search {
						count++
					}
				}
			}
		}

		if count > 0 {
			camein = true
		}
	}

	return position
}
