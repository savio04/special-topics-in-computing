package maximumvalue

func Maximumvaluev1(array []int, size int) int {
	max := array[0]

	for index := 0; index < size; index++ {
		if array[index] > max {
			max = array[index]
		}
	}

	return max
}

func max(first int, second int) int {
	if first > second {
		return first
	}

	return second
}

func Maximumvaluev2(array []int, start int, end int) int {
	if (end - start) <= 1 {
		return max(array[start], array[end])
	} else {
		middle := (start + end) / 2
		max1 := Maximumvaluev2(array, start, middle)
		max2 := Maximumvaluev2(array, middle+1, end)

		return max(max1, max2)
	}
}
