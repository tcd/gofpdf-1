package gofpdf

import "io"

// IObj inteface for all pdf objects.
type IObj interface {
	String() string
	write(w io.Writer, objID int) error
}

// ProcsetIObj interface for all pdf objects that need to be registerd to the procset
type ProcsetIObj interface {
	IObj
	procsetIdentifier() string
}
