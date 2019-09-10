package gofpdf

import (
	"fmt"
	"io"
)

// Page boundary types
const (
	PageBoundaryCrop  = iota // the region that the PDF viewer application is expected to display or print.
	PageBoundaryMedia        // the width and height of the page. For the average user, this probably equals the actual page size.
	PageBoundaryBleed        // the region to which the page contents needs to be clipped when output in a production environment. Usually the BleedBox is 3 to 5 millimeters larger than the TrimBox. By default the BleedBox equals the CropBox.
	PageBoundaryTrim         // the intended dimensions of the finished page; the actual page size that gets printed.
	PageBoundaryArt          // infrequently used; defines a page’s meaningful content area.
	pageBoundaryMax          // the end of the road
)

//PageOption option of page
type PageOption struct {
	PageBoundaries [5]*PageBoundary
}

// IsEmpty returns true if no PageBoundaries have been set.
func (po *PageOption) IsEmpty() bool {
	return len(po.PageBoundaries) == 0
}

func (gp *Fpdf) NewPageOption(w, h float64) *PageOption {
	return NewPageOption(gp.curr.unit, w, h)
}

func NewPageOption(u int, w, h float64) (po *PageOption) {
	po = &PageOption{}
	po.AddPageBoundary(NewPageSizeBoundary(u, w, h))
	return
}

func (po *PageOption) AddPageBoundary(pb *PageBoundary) {
	po.PageBoundaries[pb.Type] = pb
}

func (po *PageOption) writePageBoundaries(w io.Writer) error {
	var cpb *PageBoundary

	for x := 0; x < pageBoundaryMax; x++ {
		if po.PageBoundaries[x] != nil {
			cpb = po.PageBoundaries[x]
		}

		if cpb == nil {
			continue
		}
		// run the last thing that wasn't null
		_, err := fmt.Fprintf(w, "/%s [%.2f %.2f %.2f %.2f]\n",
			PageBoundaryType(x),
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

func (po *PageOption) GetBoundary(t int) *PageBoundary {
	for ; t >= PageBoundaryMedia; t-- {
		if po.PageBoundaries[t] != nil {
			return po.PageBoundaries[t]
		}
	}

	return nil
}

func (po PageOption) merge(po2 PageOption) PageOption {
	var pageOpt PageOption

	for x := 0; x < pageBoundaryMax; x++ {
		if po.PageBoundaries[x] != nil {
			pageOpt.AddPageBoundary(po.PageBoundaries[x])
		}

		if po2.PageBoundaries[x] != nil {
			pageOpt.AddPageBoundary(po2.PageBoundaries[x])
		}
	}

	return pageOpt
}

type PageBoundary struct {
	Type     int
	Position Point
	Size     Rect
}

func PageBoundaryType(t int) string {
	switch t {
	case PageBoundaryMedia:
		return "MediaBox"
	case PageBoundaryCrop:
		return "CropBox"
	case PageBoundaryBleed:
		return "BleedBox"
	case PageBoundaryTrim:
		return "TrimBox"
	case PageBoundaryArt:
		return "ArtBox"
	}

	return ""
}

func NewPageBoundary(u int, t int, x, y, w, h float64) (*PageBoundary, error) {
	if t >= pageBoundaryMax {
		return nil, fmt.Errorf("Page boundary %d is not valid", t)
	}

	UnitsToPointsVar(u, &x, &y, &w, &h)
	return &PageBoundary{
		Type:     t,
		Position: Point{X: x, Y: y},
		Size:     Rect{W: w, H: h},
	}, nil
}

func (gp *Fpdf) NewPageBoundary(t int, x, y, w, h float64) (*PageBoundary, error) {
	return NewPageBoundary(gp.curr.unit, t, x, y, w, h)
}

func NewPageSizeBoundary(u int, w, h float64) *PageBoundary {
	pb, _ := NewPageBoundary(u, PageBoundaryMedia, 0, 0, w, h)
	return pb
}

func (gp *Fpdf) NewPageSizeBoundary(w, h float64) *PageBoundary {
	pb, _ := gp.NewPageBoundary(PageBoundaryMedia, 0, 0, w, h)
	return pb
}

func NewCropPageBoundary(u int, x, y, w, h float64) *PageBoundary {
	pb, _ := NewPageBoundary(u, PageBoundaryCrop, x, y, w, h)
	return pb
}

func (gp *Fpdf) NewCropPageBoundary(x, y, w, h float64) *PageBoundary {
	pb, _ := gp.NewPageBoundary(PageBoundaryCrop, x, y, w, h)
	return pb
}

func NewBleedPageBoundary(u int, x, y, w, h float64) *PageBoundary {
	pb, _ := NewPageBoundary(u, PageBoundaryBleed, x, y, w, h)
	return pb
}

func (gp *Fpdf) NewBleedPageBoundary(x, y, w, h float64) *PageBoundary {
	pb, _ := gp.NewPageBoundary(PageBoundaryBleed, x, y, w, h)
	return pb
}

func NewTrimPageBoundary(u int, x, y, w, h float64) *PageBoundary {
	pb, _ := NewPageBoundary(u, PageBoundaryTrim, x, y, w, h)
	return pb
}

func (gp *Fpdf) NewTrimPageBoundary(x, y, w, h float64) *PageBoundary {
	pb, _ := gp.NewPageBoundary(PageBoundaryTrim, x, y, w, h)
	return pb
}

func NewArtPageBoundary(u int, x, y, w, h float64) *PageBoundary {
	pb, _ := NewPageBoundary(u, PageBoundaryArt, x, y, w, h)
	return pb
}

func (gp *Fpdf) NewArtPageBoundary(x, y, w, h float64) *PageBoundary {
	pb, _ := gp.NewPageBoundary(PageBoundaryArt, x, y, w, h)
	return pb
}
