package ternarysearch

func Ternarysearch(array []int, search int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		left_mid := low + (high-low)/3
		right_mid := high - (high-low)/3

		if array[left_mid] == search {
			return left_mid
		} else if array[right_mid] == search {
			return right_mid
		} else if array[left_mid] > search {
			high = left_mid - 1
		} else if array[right_mid] < search {
			low = right_mid + 1
		} else {
			low = left_mid + 1
			high = right_mid - 1
		}
	}

	return -1
}
