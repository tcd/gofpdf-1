package gofpdf

import (
	"fmt"
	"io"
	"strings"
)

const pageType = "Page"

// PageObj pdf page object
// Implements IObj.
type PageObj struct {
	ResourcesRelate   string
	pageOption        PageOption
	Links             []linkOption
	indexOfContentObj int
	getRoot           func() *Fpdf
}

func (p *PageObj) init(funcGetRoot func() *Fpdf) {
	p.getRoot = funcGetRoot
	p.Links = make([]linkOption, 0)
	p.indexOfContentObj = -1
}

func (p *PageObj) setOption(opt PageOption) {
	p.pageOption = opt
}

func (p *PageObj) setIndexOfContentObj(index int) {
	p.indexOfContentObj = index
}

func (p *PageObj) write(w io.Writer, objID int) error {
	io.WriteString(w, "<<\n")
	fmt.Fprintf(w, "  /Type /%s\n", p.getType())
	io.WriteString(w, "  /Parent 2 0 R\n")
	fmt.Fprintf(w, "  /Resources %s\n", p.ResourcesRelate)

	var err error
	gp := p.getRoot()
	if len(p.Links) > 0 {
		io.WriteString(w, "  /Annots [")
		for _, l := range p.Links {
			if l.url != "" {
				err = p.writeExternalLink(w, l, objID)
			} else {
				err = p.writeInternalLink(w, l, gp.anchors)
			}
			if err != nil {
				return err
			}
		}
		io.WriteString(w, "]\n")
	}

	fmt.Fprintf(w, "  /Contents %d 0 R\n", p.indexOfContentObj+1)
	p.pageOption.writePageBoundaries(w)
	io.WriteString(w, ">>\n")
	return nil
}

func (p *PageObj) writeExternalLink(w io.Writer, l linkOption, objID int) error {
	protection := p.getRoot().protection()
	url := l.url
	if protection != nil {
		tmp, err := rc4Cip(protection.objectkey(objID), []byte(url))
		if err != nil {
			return err
		}
		url = string(tmp)
	}
	url = strings.Replace(url, "\\", "\\\\", -1)
	url = strings.Replace(url, "(", "\\(", -1)
	url = strings.Replace(url, ")", "\\)", -1)
	url = strings.Replace(url, "\r", "\\r", -1)

	_, err := fmt.Fprintf(w, "<</Type /Annot /Subtype /Link /Rect [%.2f %.2f %.2f %.2f] /Border [0 0 0] /A <</S /URI /URI (%s)>>>>",
		l.x, l.y, l.x+l.w, l.y-l.h, url)
	return err
}

func (p *PageObj) writeInternalLink(w io.Writer, l linkOption, anchors map[string]anchorOption) error {
	a, ok := anchors[l.anchor]
	if !ok {
		return nil
	}
	_, err := fmt.Fprintf(w, "<</Type /Annot /Subtype /Link /Rect [%.2f %.2f %.2f %.2f] /Border [0 0 0] /Dest [%d 0 R /XYZ 0 %.2f null]>>",
		l.x, l.y, l.x+l.w, l.y-l.h, a.page+1, a.y)
	return err
}

func (p *PageObj) getType() string {
	return pageType
}

func (p *PageObj) getContent() *ContentObj {
	return p.getRoot().pdfObjs.getPageContent(p)
}
