package paths

import (
	"log"
	"os"
	"runtime"
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
	pathSeparator := string(os.PathSeparator)
	if runtime.GOOS == "windows" {
		file = "file:" + pathSeparator + pathSeparator
	}
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePatch := file + path
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
	pathSeparator := string(os.PathSeparator)
	var result string
	for i, e :=  range elem {
		if i == len(elem) - 1 {
			result += e
			break
		}
		result += e + pathSeparator
	}
	return result
}