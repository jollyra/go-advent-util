package main

import "testing"

// Insert inserts integer x after xs[index].
func Insert(xs []int, x, i int) []int {
	return append(xs[:i+1], append([]int{x}, xs[i+1:]...)...)
}

// Remove removes the integer after xs[index].
func Remove(xs []int, i int) []int {
	return append(xs[:i], xs[i+1:]...)
}

func TestInsert(t *testing.T) {
	cases := []struct {
		xs   []int
		x    int
		i    int
		want []int
	}{
		{
			[]int{1, 2, 3},
			9,
			1,
			[]int{1, 2, 9, 3},
		},
		{
			[]int{1},
			2,
			0,
			[]int{1, 2},
		},
	}
	for _, c := range cases {
		got := Insert(c.xs, c.x, c.i)
		if !equal(got, c.want) {
			t.Errorf("Insert(%v, %d, %d) == %v, want %v",
				c.xs, c.x, c.i, got, c.want)
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		xs   []int
		i    int
		want []int
	}{
		{
			[]int{1, 2, 3},
			1,
			[]int{1, 3},
		},
		{
			[]int{1},
			0,
			[]int{},
		},
	}
	for _, c := range cases {
		got := Remove(c.xs, c.i)
		if !equal(got, c.want) {
			t.Errorf("Remove(%v, %d) == %v, want %v",
				c.xs, c.i, got, c.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
