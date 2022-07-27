package createpdf

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

//Создаем документ с нужными параметрами
func Create(s *strings.Reader) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if  err != nil {
		log.Fatal(err)
	}
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.MarginTop.Set(0)

	page := wkhtmltopdf.NewPageReader(s)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)  

	page.Allow.Set(path +  "/templates")
	page.FooterRight.Set("[page]")

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./result/result.pdf")
	if err != nil {
		log.Fatal(err)
	}

}