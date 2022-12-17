package lib

import (
	"bufio"
	"errors"
	"os"

	"golang.org/x/exp/constraints"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(path string) []string {
	f, err := os.Open(path)
	Check(err)
	defer f.Close()

	var lines = make([]string, 0)

	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	return lines
}

func FindMax[V constraints.Ordered](slice []V) (*V, error) {
	if len(slice) == 0 {
		return nil, errors.New("Empty slice")
	}
	var max V = slice[0]
	for i, v := range slice {
		if i == 0 || v > max {
			max = v
		}
	}
	return &max, nil
}

func Sum[V constraints.Integer](slice []V) V {
	var s V = 0
	for _, n := range slice {
		s += n
	}
	return s
}

/**
 * map function for slice: maps slice[I] to slice[R] by
 * using f(T) R
 */
func Map[I any, R any](s *[]I, f func(item I) R) []R {
	var result = make([]R, 0, len(*s))
	for _, item := range *s {
		result = append(result, f(item))
	}
	return result
}

func Max[T constraints.Ordered](a T, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}

func AbsInt64(a int64) int64 {
	if a < 0 {
		return -1 * a
	}
	return a
}

func Contains[T comparable](list []T, el T) bool {
	for _, a := range list {
		if a == el {
			return true
		}
	}
	return false
}

// greatest common divisor (GCD) via Euclidean algorithm
// source: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// source: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
