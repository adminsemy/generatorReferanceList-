package paths

import (
	"log"
	"os"
	"runtime"
	"strings"
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
	var file string
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		file = "file:" + "/" + "/" + "/"
	}
	//Указываем абсолютный путь до проекта. Если Windows то
	//заменяем обратные слеши на обычные
	filePatch := file + replase(path)
	pathPitures := join(filePatch, "templates", "picture")
	pathCss := join(filePatch, "templates", "css", "main.css")
	pathTemplate := join(path, "templates", "referanceList.html")
	return &Paths{
		path,
		pathPitures,
		pathCss,
		pathTemplate,
	}
}

func join(elem ...string) string {
	result := strings.Join(elem, "/")
	return result
}

func replase(s string) string{
	result := strings.ReplaceAll(s, "\\", "/")
	return result
}