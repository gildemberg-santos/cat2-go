# cat2-go
Extrair texto de pdf

``` go
new_pdf := services.ExtractPDF{}
new_pdf.New("<File Name>.pdf")
new_pdf.Open()
// new_pdf.FileName
// new_pdf.Text

new_txt := services.ExtractTXT{}
new_txt.New("<File Name>.txt")
new_txt.Open()
// new_txt.FileName
// new_txt.Text
```
