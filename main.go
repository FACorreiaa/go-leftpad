package main

import (
	"fmt"
	"strings"
	"time"
)

func leftPad(str string, length int, ch rune) string {
	if length <= len(str) {
		return str
	}
	return strings.Repeat(string(ch), length-len(str)) + str
}

func run(fn func(string, int, rune) string, count int, args ...interface{}) time.Duration {
	start := time.Now()
	for i := 0; i < count; i++ {
		fn(args[0].(string), args[1].(int), args[2].(rune))
	}
	return time.Since(start)
}

func main() {
	for _, x := range []int{10, 100, 1000, 10000} {
		for _, y := range []int{10, 100, 1000, 10000} {
			duration := run(leftPad, 100000, "foo", x, ' ', y)
			fmt.Printf("leftPad %d %d %v\n", x, y, duration)
		}
	}
}
