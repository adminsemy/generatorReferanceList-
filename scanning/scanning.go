package scanning

import (
	"io/ioutil"
	"log"
)

//Сканируем нужную директорию для списка имен файлов в директории
func Scan(directory string) []string {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	var filesName []string
	for _, file := range files {
		filesName = append(filesName, file.Name())
	}
	return filesName
}