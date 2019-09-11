package gofpdf

// Current state
type Current struct {
	setXCount int     // Number of times SetX() has been called
	X         float64 // Horizontal position in the document
	Y         float64 // Vertical position in the document

	//font
	IndexOfFontObj int
	CountOfFont    int
	CountOfL       int

	FontSize  float64
	FontStyle int // Regular | Bold | Italic | Underline
	FontCount int
	FontType  int // CURRENT_FontType_IFONT or  CURRENT_FontType_SUBSET

	FontISubset *SubsetFontObj // FontType == CURRENT_FontType_SUBSET
	TextOption

	// page
	IndexOfPageObj int

	// img
	CountOfImg int

	// text color
	txtColor Rgb

	// text grayscale
	grayFill float64

	// draw grayscale
	grayStroke float64

	lineWidth float64
	capStyle  int
	joinStyle int

	lheight    float64
	unit       Unit
	pageOption PageOption
}

func (c *Current) setTextColor(rgb Rgb) {
	c.txtColor = rgb
}

func (c *Current) textColor() Rgb {
	return c.txtColor
}

func (c *Current) setLineHeight(h float64) float64 {
	if h <= 0 {
		c.lheight = h
	}

	return c.lheight
}

//Rgb  rgb color
type Rgb struct {
	r uint8
	g uint8
	b uint8
}

//SetR set red
func (rgb *Rgb) SetR(r uint8) {
	rgb.r = r
}

//SetG set green
func (rgb *Rgb) SetG(g uint8) {
	rgb.g = g
}

//SetB set blue
func (rgb *Rgb) SetB(b uint8) {
	rgb.b = b
}

func (rgb Rgb) equal(obj Rgb) bool {
	if rgb.r == obj.r && rgb.g == obj.g && rgb.b == obj.b {
		return true
	}
	return false
}
