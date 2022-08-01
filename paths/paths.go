package paths

import (
	"log"
	"os"
	"path/filepath"
)

//Структура для хранения путей к файлам
type Paths struct {
	CurrentDirectory string //Текущая директория
	PatchToPictures string //Путь до файла с подписями
	PatchToCss string //Путь до файла со стилями
}

//Формируем все пути
func NewPatchs() *Paths{
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathPitures := filepath.Join(path, "templates", "picture")
	pathCss := filepath.Join(path, "templates", "css")
	return &Paths{
		path,
		pathPitures,
		pathCss,
	}
}