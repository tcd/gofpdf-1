package cache

import (
	"io"

	. "github.com/tcd/gofpdf-1"
)

type iCacheContent interface {
	write(w io.Writer, protection *PDFProtection) error
}
