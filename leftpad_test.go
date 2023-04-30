package main

import (
	"fmt"
	"testing"
)

func BenchmarkLeftPad(b *testing.B) {
	for _, x := range []int{10, 100, 1000, 10000} {
		for _, y := range []int{10, 100, 1000, 10000} {
			b.Run(
				fmt.Sprintf("x=%d y=%d", x, y),
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						leftPad("foo", x, ' ')
					}
				},
			)
		}
	}
}
