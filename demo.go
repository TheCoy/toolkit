// Package golib is a personal go library
//
// It provides a list of common code snippets
package golib

import "fmt"
import "math/rand"
import "time"

// SayHello It is a say hi demo
func SayHello(lan string) {
	fmt.Println("hello,", lan)
}

// GetRandNum generates a number of integers
func GetRandNum(count int) []int {
	var numbers []int
	for i := 0; i < count; i++ {
		times := int64(time.Now().Nanosecond())
		rand.Seed(times)
		num := rand.Intn(1000)
		numbers = append(numbers, num)
	}

	return numbers
}

// SwitchDemo shows some switch characters
func SwitchDemo(k int) {
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	n, h := len(runes), len(runes)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
