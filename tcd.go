package gofpdf

// GetFontSize returns the current font size (not sure in what units).
func (gp *Fpdf) GetFontSize() float64 {
	return gp.curr.Font_Size
}
