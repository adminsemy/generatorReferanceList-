package main

import (
	"fmt"
	"log"
	"os"

	createpdf "github.com/adminsemy/generatorReferanceList/createPdf"
	htmlgenerator "github.com/adminsemy/generatorReferanceList/htmlGenerator"
	"github.com/adminsemy/generatorReferanceList/interactive"
	"github.com/adminsemy/generatorReferanceList/scanning"
	"github.com/adminsemy/generatorReferanceList/signatories"
)

func main(){
	var data htmlgenerator.Data

	//Устанавлваем текущую директорию
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	data.CurrentDirectory = path + "/templates/picture"

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
	fmt.Println(listForChoice)
	//Формируем ответы на вопросы
	answers, err := interactive.NewAnswers(listForChoice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answers)

	data.Answers = answers
	//Создаен из шаблона сгенерированную строку
	g := htmlgenerator.Generate("./templates/referanceList.html", &data)

	//Формируем новый pdf документ
	createpdf.Create(g, path)
	fmt.Println("Done!")
}