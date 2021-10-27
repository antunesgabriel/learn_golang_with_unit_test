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
