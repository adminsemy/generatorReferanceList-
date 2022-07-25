package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/adminsemy/generatorReferanceList/scan"
	"github.com/adminsemy/generatorReferanceList/signatories"
)

func main(){
	scan.Scan()
	files, err := ioutil.ReadDir("./templates/picture")
	if err != nil {
		log.Fatal(err)
	}
	var filesName []string
	for _, file := range files {
		filesName = append(filesName, file.Name())
	}
	results, err := signatories.AddSignatories(filesName)
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}


	//Формируем новый pdf документ
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