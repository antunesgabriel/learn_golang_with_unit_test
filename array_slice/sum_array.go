package arrayslice

func SumArray(data [5]int) (sum int) {
	for _, value := range data { // how to use range -> https://gobyexample.com/range
		sum += value
	}

	return
}

func SumSlice(data []int) int { // more about slice:  https://golang.org/doc/effective_go#slices
	sum := 0

	for _, value := range data {
		sum += value
	}

	return sum
}

func SumAll(data ...[]int) (sumSlice []int) {
	length := len(data)
	sumSlice = make([]int, length)

	for idx, values := range data {
		sumSlice[idx] = SumSlice(values)
	}

	return

	/**
		  Another way to do it:

		  func SumAll(data ...[]int) (sumSlice []int) {
		      for _, values := range data {
		          sumSlice = append(sumSlice, SumSlice(values))
		      }

		      return
		  }

	      Benchmark using Append:
	      BenchmarkSumAll-4        9953186               110.4 ns/op

	      Benchmarkl using make and indexing item:
	      BenchmarkSumAll-4       26162486                41.29 ns/op
	*/
}

func SumAllSliceRest(data ...[]int) (sumSlice []int) {
	length := len(data)
	sumSlice = make([]int, length)

	for idx, values := range data {

		if len(values) == 0 {
			sumSlice[idx] = 0
			continue
		}

		rest := values[1:] // more about slice https://go.dev/blog/slices-intro
		sumSlice[idx] = SumSlice(rest)
	}

	return
}
