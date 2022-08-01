package paths

import (
	"log"
	"os"
	"path/filepath"
)

//Структура для хранения путей к файлам
type Paths struct {
	CurrentDirectory string //Текущая директория
	PathToPictures string //Путь до файла с подписями
	PathToCss string //Путь до файла со стилями
	PathToTemplate string //Путь до шаблона html документа
}

//Формируем все пути
func NewPatchs() *Paths{
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathPitures := filepath.Join(path, "templates", "picture")
	pathCss := filepath.Join(path, "templates", "css", "main.css")
	pathTemplate := filepath.Join(path, "templates", "referanceList.html")
	return &Paths{
		path,
		pathPitures,
		pathCss,
		pathTemplate,
	}
}