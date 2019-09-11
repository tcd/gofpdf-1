package gofpdf

import (
	"fmt"
	"io"
)

type FontDescriptorObj struct {
	font              IFont
	fontFileObjRelate string
}

func (f *FontDescriptorObj) String() string { return "FontDescriptor" }

func (f *FontDescriptorObj) init(funcGetRoot func() *Fpdf) {}

func (f *FontDescriptorObj) write(w io.Writer, objID int) error {

	fmt.Fprintf(w, "<</Type /FontDescriptor /FontName /%s ", f.font.GetName())
	descs := f.font.GetDesc()
	i := 0
	max := len(descs)
	for i < max {
		fmt.Fprintf(w, "/%s %s ", descs[i].Key, descs[i].Val)
		i++
	}

	if f.String() == "Type1" {
		io.WriteString(w, "/FontFile ")
	} else {
		io.WriteString(w, "/FontFile2 ")
	}

	io.WriteString(w, f.fontFileObjRelate)
	io.WriteString(w, ">>\n")

	return nil
}

func (f *FontDescriptorObj) SetFont(font IFont) {
	f.font = font
}

func (f *FontDescriptorObj) GetFont() IFont {
	return f.font
}

func (f *FontDescriptorObj) SetFontFileObjRelate(relate string) {
	f.fontFileObjRelate = relate
}
