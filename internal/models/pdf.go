package models

import (
	"io"

	"github.com/dslipak/pdf"
)

type ModelPDF struct {
	FileName string
	Text     string
}

func (p *ModelPDF) New(filename string) {
	p.FileName = filename
}

func (p *ModelPDF) Open() {
	content, err := pdf.Open(p.FileName)
	if err != nil {
		panic(err)
	}

	reader, err := content.GetPlainText()
	if err != nil {
		panic(err)
	}

	reader_bytes, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	p.Text = string(reader_bytes)

}
