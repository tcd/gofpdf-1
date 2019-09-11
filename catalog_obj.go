package gofpdf

import (
	"fmt"
	"io"
)

// CatalogObj catalog dictionary.
// Implements the IObj interface.
type CatalogObj struct{}

func (c *CatalogObj) String() string {
	return "Catalog"
}

func (c *CatalogObj) init(funcGetRoot func() *Fpdf) {}

func (c *CatalogObj) write(w io.Writer, objID int) error {
	io.WriteString(w, "<<\n")
	fmt.Fprintf(w, "  /Type /%s\n", c)
	io.WriteString(w, "  /Pages 2 0 R\n")
	io.WriteString(w, ">>\n")
	return nil
}
