package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main(){
	files, err := ioutil.ReadDir("./templates/picture")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if  err != nil {
		log.Fatal(err)
	}
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	page := wkhtmltopdf.NewPage("./templates/referanceList.html")

	page.Allow.Set("./templates")
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

	fmt.Println("Done!")
}