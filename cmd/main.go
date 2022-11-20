package main

import (
	"fmt"
	"legal-research-automation-go/pkg/services"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No arguments provided")
		return
	}

	filename := os.Args[1:][0]
	filename_extension := filename[len(filename)-3:]

	if filename_extension == "pdf" {
		pdf := services.ExtractPDF{}
		pdf.New(filename)
		pdf.Open()
		fmt.Println(pdf.Text)
		return
	}

	txt := services.ExtractTXT{}
	txt.New(filename)
	txt.Open()
	fmt.Println(txt.Text)
}
