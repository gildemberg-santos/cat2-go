package services

import (
	"os"
	"strings"
)

type ExtractTXT struct {
	FileName string
	Text     string
}

func (p *ExtractTXT) New(filename string) {
	p.FileName = filename
}

func (p *ExtractTXT) Open() {
	reader_bytes, err := os.ReadFile(p.FileName)
	if err != nil {
		panic(err)
	}
	p.Text = strings.Replace(string(reader_bytes), "\n", "", -1)
}
