package arrayslice

func SumArray(data [5]int) int {
	sum := 0

	for _, value := range data { // how to use range -> https://gobyexample.com/range
		sum += value
	}

	return sum
}

func SumSlice(data []int) int { // more about slice:  https://golang.org/doc/effective_go#slices
	sum := 0

	for _, value := range data {
		sum += value
	}

	return sum
}

func SumAll(data ...[]int) (sumSlice []int) {
	lenght := len(data)
	sumSlice = make([]int, lenght)

	for idx, values := range data {
		sumSlice[idx] = SumSlice(values)
	}

	return
}
