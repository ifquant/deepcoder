package dsl

import (
	"sort"
)

// First-order functions

// Head returns its first element of an slice.
func Head(s []int) (ret int, ok bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return 0, false
}

// Last returns its last element of an slice.
func Last(s []int) (ret int, ok bool) {
	if len(s) > 0 {
		return s[len(s)-1], true
	}
	return 0, false
}

// Take returns the slice truncated after the n-th element. (If the length of s was no larger than n in the first place, it is returned without modification.)
func Take(s []int, n int) (ret []int) {
	if n > len(s) {
		return s
	}
	return s[:n]
}

// Drop returns the slice with the first n elements dropped. (If the length of s was no larger than n in the first place, an empty array is returned.)
func Drop(s []int, n int) (ret []int) {
	if n > len(s) {
		return nil
	}
	return s[n:]
}

// Access returns the (i+1)th element of slice s.
func Access(s []int, i int) (ret int, ok bool) {
	if i >= len(s) {
		return 0, false
	}
	return s[i], true
}

func Minimun(s []int) (ret int, ok bool) {
	if len(s) > 0 {
		sort.Ints(s)
		return s[len(s)-1], true
	}
	return 0, false
}

func Maximun(s []int) (ret int, ok bool) {
	if len(s) > 0 {
		sort.Ints(s)
		return s[0], true
	}
	return 0, false
}
func Reverse(s []int) (ret []int) {
	for i := 0; i < len(s)-1; i++ {
		s[i], s[len(s)-i] = s[len(s)-i], s[i]
	}
	return s
}

func Sort(s []int) (ret []int) {
	sort.Ints(s)
	return s
}

func Sum(s []int) (ret int) {
	for _, i := range s {
		ret += i
	}
	return
}

// Higher-order functions
type F func(int) int

func Map(s []int, f F) []int {
	for i := range s {
		s[i] = f(s[i])
	}
	return s
}

type B func(int) bool

func Filter(s []int, b B) []int {
	var n []int
	for _, i := range s {
		if b(i) {
			n = append(n, i)
		}
	}
	return n

}

func Count(s []int, b B) (ret int) {
	for _, i := range s {
		if b(i) {
			ret += i
		}
	}
	return
}

type Z func(x, y int) int

func ZipWith(xs, ys []int, z Z) []int {
	l := len(xs)
	if len(ys) < len(xs) {
		l = len(ys)
	}
	var n []int
	for i := 0; i < l; i++ {
		n = append(n, z(xs[i], ys[i]))
	}
	return n
}

func Scanl1(xs []int, z Z) []int {
	if len(xs) == 1 {
		return xs
	}
	var n = make([]int, 1, len(xs))
	if len(xs) > 1 {
		n[0] = xs[0]
		for i := 1; i < len(xs); i++ {
			n[i] = z(n[i-1], xs[i])
		}

	}
	return n
}

// self-explanatory lambdas, TODO
func int2Int(s string) func(int) int {
	return nil
}

func int2Bool(s string) func(int) bool {
	return nil
}

func int2Int2Int(s string) func(int, int) int {
	return nil
}
