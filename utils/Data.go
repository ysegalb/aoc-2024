package utils

type Point struct {
	X int
	Y int
}

// Up returns a new Point that is above the receiver (X, Y-1).
func (p *Point) Up() Point {
	return Point{X: p.X, Y: p.Y - 1}
}

// Down returns a new Point that is below the receiver (X, Y+1).
func (p *Point) Down() Point {
	return Point{X: p.X, Y: p.Y + 1}
}

// Left returns a new Point that is to the left of the receiver (X-1, Y).
func (p *Point) Left() Point {
	return Point{X: p.X - 1, Y: p.Y}
}

// Right returns a new Point that is to the right of the receiver (X+1, Y).
func (p *Point) Right() Point {
	return Point{X: p.X + 1, Y: p.Y}
}

func (p *Point) UpLeft() Point {
	return Point{X: p.X - 1, Y: p.Y - 1}
}

func (p *Point) UpRight() Point {
	return Point{X: p.X + 1, Y: p.Y - 1}
}

func (p *Point) DownLeft() Point {
	return Point{X: p.X - 1, Y: p.Y + 1}
}

func (p *Point) DownRight() Point {
	return Point{X: p.X + 1, Y: p.Y + 1}
}

type Node[T comparable] struct {
	Content T
	North   *Node[T]
	East    *Node[T]
	South   *Node[T]
	West    *Node[T]
}

func (n *Node[T]) ForEach(f func(*Node[T])) {
	for _, dir := range []*Node[T]{n.North, n.East, n.South, n.West} {
		if dir != nil {
			Apply(dir, f)
		}
	}
}

func (n *Node[T]) AllMatch(f func(*Node[T]) bool) bool {
	for _, dir := range []*Node[T]{n.North, n.East, n.South, n.West} {
		if dir != nil {
			if !f(dir) {
				return false
			}
		}
	}
	return true
}

type Row[T comparable] []T

type Grid[T comparable] struct {
	Rows []Row[T]
}

func (g *Grid[T]) SamePoints(one, two Point) bool {
	return g.GetPoint(one) == g.GetPoint(two)
}

func (g *Grid[T]) Height() int {
	return len(g.Rows)
}

func (g *Grid[T]) Width() int {
	return len(g.Rows[0])
}

func (g *Grid[T]) GetPoint(p Point) T {
	return g.Get(p.X, p.Y)
}

func (g *Grid[T]) All() func(yield func(int, []T) bool) {
	return func(yield func(int, []T) bool) {
		// Lógica de iteración
		for i := 0; i < len(g.Rows); i++ {
			if !yield(i, g.Rows[i]) {
				return
			}
		}
	}
}

// NewGrid returns a new, empty Grid[T].
func NewGrid[T comparable]() Grid[T] {
	return Grid[T]{Rows: make([]Row[T], 0)}
}

// Get returns the element at the given coordinates.
//
// The element at the given coordinates is returned. If the given coordinates
// are out of bounds, the behavior is undefined.
func (g *Grid[T]) Get(x, y int) T {
	return g.Rows[y][x]
}

// Append adds a new row to the grid with a single element, the given value.
func (g *Grid[T]) Append(value T) {
	g.Rows = append(g.Rows, []T{value})
}

// AppendAll appends the given values to the grid.
//
// The given values are appended as a new row to the grid.
func (g *Grid[T]) AppendAll(values []T) {
	g.Rows = append(g.Rows, append([]T{}, values...))
}

// Apply calls the given function `f` with the given value `t` as argument.
//
// This is a simple convenience function to avoid having to write an
// anonymous function just to call `f` with `t`:
//
//	Apply(x, f)  // Instead of:
//	f(x)
func Apply[T any](t T, f func(T)) {
	f(t)
}
