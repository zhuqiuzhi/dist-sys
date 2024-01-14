package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Tree[T interface{}] struct {
	left, right *Tree[T]
	a           []T
	value       T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] { return nil }

// Scale returns a copy of s with each element multiplied by c.
// This implementation has a problem, as we will see.
func Scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

func main() {
	var a, b, m float64

	m = GMin(a, b) // no type argument
	fmt.Println(m)

	// var stringTree Tree[string]
}
