package cache

import (
	"io"

	. "github.com/tcd/gofpdf-1"
)

type cacheContentClosePath struct {
}

func (c *cacheContentClosePath) write(w io.Writer, protection *PDFProtection) error {
	io.WriteString(w, "h\n")
	return nil
}
