package arrayslice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumArray(t *testing.T) {
	feedback := func(t *testing.T, expected *int, result *int, data *[5]int) {
		if *expected != *result {
			t.Errorf("[FAILED] Expected: %d - Result: %d - Data: %v", *expected, *result, *data)
		}
	}

	t.Run("i should sum array int", func(t *testing.T) {
		data := [5]int{1, 2, 3, 4, 5}
		expected := 15
		result := SumArray(data)

		feedback(t, &expected, &result, &data)
	})

}

func ExampleSumArray() {
	data := [...]int{10, 10, 10, 10, 10}
	result := SumArray(data)

	fmt.Println(result)
	// Output: 50
}

func BenchmarkSumArray(b *testing.B) {
	data := [...]int{10, 10, 10, 10, 10}

	for i := 0; i < b.N; i++ {
		SumArray(data)
	}
}

func TestSumSlice(t *testing.T) {
	feedback := func(t *testing.T, expected *int, result *int, data *[]int) {
		if *expected != *result {
			t.Errorf("[FAILED] Expected: %d - Result: %d - Data %v", *expected, *result, *data)
		}
	}

	t.Run("i should sum values into slice", func(t *testing.T) {
		data := []int{5, 5, 45}
		expected := 55
		result := SumSlice(data)

		feedback(t, &expected, &result, &data)
	})
}

func ExampleSumSlice() {
	data := []int{14, 14, 14, 14, 14, 14, 14}
	result := SumSlice(data)

	fmt.Println(result)
	// Output: 98
}

func BenchmarkSumSlice(b *testing.B) {
	data := []int{4, 5, 6, 7, 8, 9, 2, 4}

	for i := 0; i < b.N; i++ {
		SumSlice(data)
	}
}

func TestSumDataSlice(t *testing.T) {
	feedback := func(expected *[]int, result *[]int, data ...[]int) {
		if !reflect.DeepEqual(*expected, *result) {
			t.Errorf("[FAILED] - Expected: %v - Result: %v - Data: %v", *expected, *result, data)
		}
	}

	t.Run("sum data slices", func(t *testing.T) {
		data1 := []int{3, 3, 3}
		data2 := []int{5, 5, 5, 5}

		expected := []int{9, 20}
		result := SumDataSlice(data1, data2)

		feedback(&expected, &result, data1, data2)
	})
}

func ExampleSumDataSlice() {
	data1 := []int{7, 7}
	data2 := []int{4, 4, 4}

	result := SumDataSlice(data1, data2)

	fmt.Println(result)
	// Output: [14 12]
}

func BenchmarkSumDataSlice(b *testing.B) {
	data1 := []int{4, 4, 4}
	data2 := []int{8, 8, 8, 8}
	data3 := []int{18, 17, 2, 45}

	for i := 0; i < b.N; i++ {
		SumDataSlice(data1, data2, data3)
	}
}
