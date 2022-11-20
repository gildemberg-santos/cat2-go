package services_test

import (
	"legal-research-automation-go/pkg/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTXT(t *testing.T) {
	filename := "Test TXT File.txt"
	txtconten := "Hello, World 2!"
	new_pdf := services.ExtractTXT{}
	new_pdf.New(filename)
	new_pdf.Open()
	assert.Equal(t, filename, new_pdf.FileName)
	assert.Equal(t, txtconten, new_pdf.Text)
}
