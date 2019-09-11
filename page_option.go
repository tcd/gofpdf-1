package gofpdf

import (
	"fmt"
	"io"
)

// BoundaryBox values represent different PDF page boundries.
type BoundaryBox int

// BoundaryBox types.
const (
	CropBox  BoundaryBox = iota // the region that the PDF viewer application is expected to display or print.
	MediaBox                    // the width and height of the page. For the average user, this probably equals the actual page size.
	BleedBox                    // the region to which the page contents needs to be clipped when output in a production environment. Usually the BleedBox is 3 to 5 millimeters larger than the TrimBox. By default the BleedBox equals the CropBox.
	TrimBox                     // the intended dimensions of the finished page; the actual page size that gets printed.
	ArtBox                      // infrequently used; defines a page’s meaningful content area.
)

// BoundaryBoxTypes can be ranged through.
var BoundaryBoxTypes = []BoundaryBox{CropBox, MediaBox, BleedBox, TrimBox, ArtBox}

// Int representation of a BoundaryBox value.
func (bb BoundaryBox) Int() int {
	return int(bb)
}

// String representation of a BoundaryBox value.
func (bb BoundaryBox) String() string {
	switch bb {
	case CropBox:
		return "CropBox"
	case MediaBox:
		return "MediaBox"
	case BleedBox:
		return "BleedBox"
	case TrimBox:
		return "TrimBox"
	case ArtBox:
		return "ArtBox"
	default:
		return ""
	}
}

// PageOption holds the BoundaryBox values of a PDF.
type PageOption struct {
	PageBoundaries [5]*PageBoundary
}

// IsEmpty returns true if no PageBoundaries have been set.
func (po *PageOption) IsEmpty() bool {
	return len(po.PageBoundaries) == 0
}

func NewPageOption(u Unit, w, h float64) (po *PageOption) {
	po = &PageOption{}
	po.AddPageBoundary(NewPageSizeBoundary(u, w, h))
	return
}

// AddPageBoundary to a PageOption.
func (po *PageOption) AddPageBoundary(pb *PageBoundary) {
	po.PageBoundaries[pb.Type] = pb
}

func (po *PageOption) writePageBoundaries(w io.Writer) error {
	var cpb *PageBoundary

	for _, box := range BoundaryBoxTypes {
		if po.PageBoundaries[box] != nil {
			cpb = po.PageBoundaries[box]
		}

		if cpb == nil {
			continue
		}
		// run the last thing that wasn't null
		_, err := fmt.Fprintf(w, "/%s [%.2f %.2f %.2f %.2f]\n",
			box.String(),
			cpb.Position.X,
			cpb.Position.Y+cpb.Size.H,
			cpb.Size.W+cpb.Position.X,
			cpb.Position.Y,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

// GetBoundary object of a given type.
func (po *PageOption) GetBoundary(t BoundaryBox) *PageBoundary {
	for ; t >= MediaBox; t-- {
		if po.PageBoundaries[t] != nil {
			return po.PageBoundaries[t]
		}
	}
	return nil
}

func (po PageOption) merge(po2 PageOption) PageOption {
	var pageOpt PageOption

	for _, box := range BoundaryBoxTypes {
		if po.PageBoundaries[box] != nil {
			pageOpt.AddPageBoundary(po.PageBoundaries[box])
		}

		if po2.PageBoundaries[box] != nil {
			pageOpt.AddPageBoundary(po2.PageBoundaries[box])
		}
	}

	return pageOpt
}

// PageBoundary is a BoundaryBox with attatched values.
type PageBoundary struct {
	Type     BoundaryBox
	Position Point
	Size     Rect
}

// NewPageBoundary returns a new PageBoundary with the given BoundaryBox type and values converted the given Unit type.
func NewPageBoundary(u Unit, t BoundaryBox, x, y, w, h float64) *PageBoundary {
	UnitsToPointsVar(u, &x, &y, &w, &h)
	return &PageBoundary{
		Type:     t,
		Position: Point{X: x, Y: y},
		Size:     Rect{W: w, H: h},
	}
}

// NewPageBoundary returns a new PageBoundary the given BoundaryBox type and values in an Fpdf's current Unit type.
func (gp *Fpdf) NewPageBoundary(t BoundaryBox, x, y, w, h float64) *PageBoundary {
	return NewPageBoundary(gp.curr.unit, t, x, y, w, h)
}

// NewPageSizeBoundary returns a new PageBoundary of the given type with given values.
func NewPageSizeBoundary(u Unit, w, h float64) *PageBoundary {
	return NewPageBoundary(u, MediaBox, 0, 0, w, h)
}

// NewPageSizeBoundary returns a new PageBoundary of the given type with given values.
func (gp *Fpdf) NewPageSizeBoundary(w, h float64) *PageBoundary {
	return gp.NewPageBoundary(MediaBox, 0, 0, w, h)
}

// NewCropPageBoundary returns a CropBox PageBoundary  with the given values.
func NewCropPageBoundary(u Unit, x, y, w, h float64) *PageBoundary {
	return NewPageBoundary(u, CropBox, x, y, w, h)
}

// NewCropPageBoundary returns a CropBox PageBoundary with the given values.
func (gp *Fpdf) NewCropPageBoundary(x, y, w, h float64) *PageBoundary {
	return gp.NewPageBoundary(CropBox, x, y, w, h)
}

// NewBleedPageBoundary returns a BleedBox PageBoundary with the given values.
func NewBleedPageBoundary(u Unit, x, y, w, h float64) *PageBoundary {
	return NewPageBoundary(u, BleedBox, x, y, w, h)
}

// NewBleedPageBoundary returns a BleedBox PageBoundary with the given values.
func (gp *Fpdf) NewBleedPageBoundary(x, y, w, h float64) *PageBoundary {
	return gp.NewPageBoundary(BleedBox, x, y, w, h)
}

// NewTrimPageBoundary returns a TrimBox PageBoundary with the given values.
func NewTrimPageBoundary(u Unit, x, y, w, h float64) *PageBoundary {
	return NewPageBoundary(u, TrimBox, x, y, w, h)
}

// NewTrimPageBoundary returns a TrimBox PageBoundary with the given values.
func (gp *Fpdf) NewTrimPageBoundary(x, y, w, h float64) *PageBoundary {
	return gp.NewPageBoundary(TrimBox, x, y, w, h)
}

// NewArtPageBoundary returns a ArtBox PageBoundary with the given values.
func NewArtPageBoundary(u Unit, x, y, w, h float64) *PageBoundary {
	return NewPageBoundary(u, ArtBox, x, y, w, h)
}

// NewArtPageBoundary returns a ArtBox PageBoundary with the given values.
func (gp *Fpdf) NewArtPageBoundary(x, y, w, h float64) *PageBoundary {
	return gp.NewPageBoundary(ArtBox, x, y, w, h)
}
