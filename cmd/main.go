package main

import "legal-research-automation-go/internal/services"

func main() {
	new_pdf := services.ModelPDF{}
	new_pdf.New("DJ SE 5869 28-07-2022.pdf")
	new_pdf.Open()
}
