package renderer

import "github.com/jroimartin/gocui"

type Point struct{ X, Y int }

type Vec struct{ X, Y float64 }

type Dim struct {
	Origin, Size Point
}

func DimFromGui(g *gocui.Gui) Dim {
	x, y := g.Size()
	return Dim{
		Origin: Point{X: 0, Y: 0},
		Size:   Point{X: x, Y: y},
	}
}

// HalfCtr returns a Dim of half the size, with the same center.
func (d Dim) HalfCtr() Dim {
	half := d.Half()
	return half.Move(half.Size.Scale(Vec{X: 0.5, Y: 0.5}))
}

func (d Dim) Half() Dim {
	return d.Scale(Vec{X: 0.5, Y: 0.5})
}

func (d Dim) Move(by Point) Dim {
	return Dim{
		Origin: d.Origin.Add(by),
		Size:   d.Size,
	}
}

func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p Point) Scale(by Vec) Point {
	return Point{
		X: int(float64(p.X) * by.X),
		Y: int(float64(p.Y) * by.Y),
	}
}

func (d Dim) Scale(v Vec) Dim {
	return Dim{
		Origin: d.Origin.Scale(v),
		Size:   d.Size.Scale(v),
	}
}
