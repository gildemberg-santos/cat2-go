package main

import "legal-research-automation-go/internal/services"

func main() {
	new_pdf := services.ModelPDF{}
	new_pdf.New("<File Name>.pdf")
	new_pdf.Open()
}
