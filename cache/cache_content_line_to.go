package cache

import (
	"fmt"
	"io"

	. "github.com/tcd/gofpdf-1"
)

type cacheContentLineTo struct {
	pageHeight float64
	x          float64
	y          float64
}

func (c *cacheContentLineTo) write(w io.Writer, protection *PDFProtection) error {
	fmt.Fprintf(w, "%.2f %.2f l\n", c.x, c.pageHeight-c.y)
	return nil
}
