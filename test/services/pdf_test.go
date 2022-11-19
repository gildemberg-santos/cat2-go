package services_test

import (
	"legal-research-automation-go/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	new_pdf := services.ModelPDF{}
	new_pdf.New("Test PDF File.pdf")
	new_pdf.Open()
	assert.Equal(t, "Test PDF File.pdf", new_pdf.FileName)
	assert.Equal(t, "Hello, World!", new_pdf.Text)
}
