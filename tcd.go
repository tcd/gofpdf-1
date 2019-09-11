package gofpdf

// GetFontSize returns the current font size (not sure in what units).
// GetFontSize returns the size of the current font in points followed by the size in the unit
func (gp *Fpdf) GetFontSize() (float64, float64) {
	return gp.curr.FontSize, PointsToUnits(gp.curr.unit, gp.curr.FontSize)
}

// GetLineHeight probably won't work like I think it will.
func (gp *Fpdf) GetLineHeight() float64 {
	return PointsToUnits(gp.curr.unit, gp.curr.FontSize)
}
