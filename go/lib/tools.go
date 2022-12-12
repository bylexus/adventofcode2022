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
