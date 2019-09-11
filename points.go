package gofpdf

// Points slice.
type Points []Point

// ToUnits will convert the Points, assumed to be in pts, to Units.
func (pts Points) ToUnits(t Unit) Points {
	points := make(Points, len(pts))

	for x := 0; x < len(pts); x++ {
		points[x] = pts[x].ToUnits(t)
	}

	return points
}

// ToPoints will convert the Points, assumed to be in units, to pts.
func (pts Points) ToPoints(t Unit) Points {
	points := make(Points, len(pts))

	for x := 0; x < len(pts); x++ {
		points[x] = pts[x].ToPoints(t)
	}

	return points
}

// Point fields X and Y specify the horizontal and vertical coordinates of
// a point, typically used in drawing.
type Point struct {
	X, Y float64
}

// XY gets the X and Y points
func (p Point) XY() (float64, float64) {
	return p.X, p.Y
}

// ToUnits will convert the point, assumed to be in pts, to the specified units.
func (p Point) ToUnits(t Unit) Point {
	return Point{
		X: PointsToUnits(t, p.X),
		Y: PointsToUnits(t, p.Y),
	}
}

// ToPoints converts the Point, assumed to be in units, to points.
func (p Point) ToPoints(t Unit) Point {
	return Point{
		X: UnitsToPoints(t, p.X),
		Y: UnitsToPoints(t, p.Y),
	}
}
