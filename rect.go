package gofpdf

// Rect is Basic 2d rectangle object for defining objects sizes.
type Rect struct {
	W            float64 // Width
	H            float64 // Height
	unitOverride int
}

// PointsToUnits converts the rectanlges width and height to Units.
// When this is called it is assumed the values of the rectangle are in Points
	if rect.unitOverride != UnitUnset {
		t = rect.unitOverride
	}

	r = Rect{W: rect.W, H: rect.H}
	PointsToUnitsVar(t, &r.W, &r.H)
	return
}

// UnitsToPoints converts the rectanlges width and height to Points.
// When this is called it is assumed the values of the rectangle are in Units
	if rect.unitOverride != UnitUnset {
		t = rect.unitOverride
	}

	r = Rect{W: rect.W, H: rect.H}
	UnitsToPointsVar(t, &r.W, &r.H)
	return
}
