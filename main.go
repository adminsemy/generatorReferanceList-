package main

import (
	"fmt"
	"log"

	createpdf "github.com/adminsemy/generatorReferanceList/createPdf"
	htmlgenerator "github.com/adminsemy/generatorReferanceList/htmlGenerator"
	"github.com/adminsemy/generatorReferanceList/interactive"
	"github.com/adminsemy/generatorReferanceList/paths"
	"github.com/adminsemy/generatorReferanceList/scanning"
	"github.com/adminsemy/generatorReferanceList/signatories"
)

func main(){
	var data htmlgenerator.Data

	//Устанавлваем текущую директорию
	path := paths.NewPatchs()
	data.Paths = path

	//Сканируем директорию с подписями 
	filesName := scanning.Scan("./templates/picture")
	results, err := signatories.AddSignatories(filesName)
	if err != nil {
		log.Fatal(err)
	}
	data.Signatories = results
	//Формируем список ознакомившихся для ответов
	var listForChoice []string
	for _, result := range data.Signatories {
		listForChoice = append(listForChoice, result.ListForChoice())
	}
	//Формируем ответы на вопросы
	answers, err := interactive.NewAnswers(listForChoice)
	if err != nil {
		log.Fatal(err)
	}

	data.Answers = answers
	//Создаен из шаблона сгенерированную строку
	g := htmlgenerator.Generate(&data)

	//Формируем новый pdf документ
	
	createpdf.Create(g, path)
	fmt.Println("Done!")
}