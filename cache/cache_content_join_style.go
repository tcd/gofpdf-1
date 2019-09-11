package cache

import (
	"fmt"
	"io"

	. "github.com/tcd/gofpdf-1"
)

type cacheContentJoinStyle struct {
	style int
}

func (c *cacheContentJoinStyle) write(w io.Writer, protection *PDFProtection) error {
	fmt.Fprintf(w, "%d j\n", c.style)
	return nil
}
