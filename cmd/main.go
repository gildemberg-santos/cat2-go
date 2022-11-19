package main

import (
	"legal-research-automation-go/internal/models"
)

func main() {
	new_pdf := models.ModelPDF{}
	new_pdf.New("DJ SE 5869 28-07-2022.pdf")
	new_pdf.Open()
}
