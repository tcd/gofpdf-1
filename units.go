package gofpdf

// Unit used for measurement in a pdf.
type Unit int

// The units that can be used in the document.
const (
	UnitUnset Unit = iota // No units were set, when conversion is called on nothing will happen
	UnitPT                // Points
	UnitMM                // Millimeters
	UnitCM                // Centimeters
	UnitIN                // Inches
	unitMax               // used to validate constant values
)

// The math needed to convert units to points.
const (
	conversionPT = 1.0
	conversionMM = 72.0 / 25.4
	conversionCM = 72.0 / 2.54
	conversionIN = 72.0
)

// UnitsToPoints converts units of the provided type to points.
func UnitsToPoints(t Unit, u float64) float64 {
	switch t {
	case UnitPT:
		return u * conversionPT
	case UnitMM:
		return u * conversionMM
	case UnitCM:
		return u * conversionCM
	case UnitIN:
		return u * conversionIN
	default:
		return u
	}
}

// PointsToUnits converts points to the provided units.
func PointsToUnits(t Unit, u float64) float64 {
	switch t {
	case UnitPT:
		return u / conversionPT
	case UnitMM:
		return u / conversionMM
	case UnitCM:
		return u / conversionCM
	case UnitIN:
		return u / conversionIN
	default:
		return u
	}
}

// UnitsToPointsVar converts units of the provided type to points for all variables supplied.
func UnitsToPointsVar(t Unit, u ...*float64) {
	for x := 0; x < len(u); x++ {
		*u[x] = UnitsToPoints(t, *u[x])
	}
}

// PointsToUnitsVar converts points to the provided units for all variables supplied.
func PointsToUnitsVar(t Unit, u ...*float64) {
	for x := 0; x < len(u); x++ {
		*u[x] = PointsToUnits(t, *u[x])
	}
}
