package iteration

import "strings"

func Repeat(word string, count int) string {
	repeatedWord := strings.Repeat(word, count)

	// for i := 0; i < count; i++ {
	// 	repeatedWord += word
	// }

	return repeatedWord
}
