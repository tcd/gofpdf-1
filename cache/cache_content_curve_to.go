package cache

import (
	"fmt"
	"io"

	. "github.com/tcd/gofpdf-1"
)

type cacheContentCurveBezierCubic struct {
	pageHeight               float64
	cx0, cy0, cx1, cy1, x, y float64
}

func (c *cacheContentCurveBezierCubic) write(w io.Writer, protection *PDFProtection) error {
	fmt.Fprintf(w, "%.5f %.5f %.5f %.5f %.5f %.5f c\n", c.cx0, (c.pageHeight - c.cy0), c.cx1, (c.pageHeight - c.cy1), c.x, (c.pageHeight - c.y))
	return nil
}
