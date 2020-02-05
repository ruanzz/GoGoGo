package main

import (
	"testing"
)

func BenchmarkTriangle(b *testing.B) {
	data := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{100, 200, 300},
	}
	b.ResetTimer()

	for i := 0; i < 10000000000; i++ {
		var ele = data[i%3]
		calTriangle(ele.a, ele.b)
	}
}

func TestTriangle(t *testing.T) {
	data := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{100, 200, 300},
	}

	for _, ele := range data {
		if result := calTriangle(ele.a, ele.b); ele.c != result {
			t.Errorf("calTriangle(%d, %d) got %d,expect: %d", ele.a, ele.b, result, ele.c)
		}
	}
}
