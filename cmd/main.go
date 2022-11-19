package main

import (
	"fmt"
	"legal-research-automation-go/pkg/services"
	"os"
)

func main() {
	args := os.Args[1:]
	new_pdf := services.ExtractPDF{}
	new_pdf.New(args[0])
	new_pdf.Open()
	fmt.Println(new_pdf.Text)
}
