package arrayslice

func SumArray(data [5]int) int {
	sum := 0

	for _, value := range data { // how to use range -> https://gobyexample.com/range
		sum += value
	}

	return sum
}

func SumSlice(data []int) int {
	sum := 0

	for _, value := range data {
		sum += value
	}

	return sum
}
