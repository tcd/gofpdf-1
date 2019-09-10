package gofpdf

// GetFontSize returns the current font size (not sure in what units).
// GetFontSize returns the size of the current font in points followed by the size in the unit
func (gp *Fpdf) GetFontSize() (float64, float64) {
	return gp.curr.Font_Size, PointsToUnits(int(gp.curr.Font_Size), Unit_MM)
}

// GetLineHeight ...
func (gp *Fpdf) GetLineHeight() float64 {
	return PointsToUnits(int(gp.curr.Font_Size), Unit_MM)

}
