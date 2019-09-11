package gofpdf

import "io"

// WriteUInt32 to an io.Writer.
func WriteUInt32(w io.Writer, v uint) error {
	a := byte(v >> 24)
	b := byte(v >> 16)
	c := byte(v >> 8)
	d := byte(v)
	_, err := w.Write([]byte{a, b, c, d})
	if err != nil {
		return err
	}
	return nil
}

// WriteUInt16 to an io.Writer.
func WriteUInt16(w io.Writer, v uint) error {
	a := byte(v >> 8)
	b := byte(v)
	_, err := w.Write([]byte{a, b})
	if err != nil {
		return err
	}
	return nil
}

// WriteTag writes a string value to an io.Writer.
func WriteTag(w io.Writer, tag string) error {
	b := []byte(tag)
	_, err := w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// WriteBytes to an io.Writer.
func WriteBytes(w io.Writer, data []byte, offset int, count int) error {
	_, err := w.Write(data[offset : offset+count])
	if err != nil {
		return err
	}
	return nil
}
