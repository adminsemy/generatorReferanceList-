package main

import (
	"fmt"
	"log"

	createpdf "github.com/adminsemy/generatorReferanceList/createPdf"
	htmlgenerator "github.com/adminsemy/generatorReferanceList/htmlGenerator"
	"github.com/adminsemy/generatorReferanceList/interactive"
	"github.com/adminsemy/generatorReferanceList/scanning"
	"github.com/adminsemy/generatorReferanceList/signatories"
)

func main(){
	//Сканируем директорию с подписями 
	filesName := scanning.Scan("./templates/picture")
	results, err := signatories.AddSignatories(filesName)
	if err != nil {
		log.Fatal(err)
	}
	//Формируем список ознакомившихся для ответов
	var listForChoice []string
	for _, result := range results {
		listForChoice = append(listForChoice, result.ListForChoice())
	}
	fmt.Println(listForChoice)
	//Формируем ответы на вопросы
	answers, err := interactive.NewAnswers(listForChoice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answers)

	//Формируем новый pdf документ
	g := htmlgenerator.Generate("./templates/referanceList.html", results, answers)
	createpdf.Create(g)
	fmt.Println("Done!")
}