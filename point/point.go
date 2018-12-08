package point

import (
	"fmt"
)

// Point represents a simple 2D point.
type Point struct{ X, Y int }

func (p Point) String() string { return fmt.Sprintf("Point{x=%d, y=%d}", p.X, p.Y) }

// Equals returns true if p and q are the same.
func (p *Point) Equals(q Point) bool {
	if p.X == q.X && p.Y == q.Y {
		return true
	}
	return false
}

// Neighbours4 returns all neighbour of point (x, y) including diagonals.
func (p Point) Neighbours4() []Point {
	points := []Point{
		Point{p.X + 1, p.Y},
		Point{p.X, p.Y + 1},
		Point{p.X - 1, p.Y},
		Point{p.X, p.Y - 1},
	}
	return points
}

// Contains returns true if q is in ps, otherwise false.
func Contains(ps []Point, q Point) bool {
	for _, p := range ps {
		if p.Equals(q) {
			return true
		}
	}
	return false
}

// ManhattanDistance return the Manhattan Distance between points a and b.
func ManhattanDistance(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
