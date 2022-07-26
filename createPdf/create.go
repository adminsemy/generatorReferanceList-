package createpdf

import (
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/adminsemy/generatorReferanceList/paths"
)

//Создаем документ с нужными параметрами
func Create(s *strings.Reader, path *paths.Paths) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if  err != nil {
		log.Fatal(err)
	}
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.MarginTop.Set(0)

	page := wkhtmltopdf.NewPageReader(s)

	page.Allow.Set(path.CurrentDirectory)

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