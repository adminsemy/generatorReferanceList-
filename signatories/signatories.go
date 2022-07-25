package signatories

import (
	"log"
	"strconv"
	"strings"
)

//Объявляем структуру, кого необходимо добавить
//в лист ознакомления. Формируем из списка файлов в
//директории templates/picture
type Signatory struct {
	SerialNumber int
	JobTitle string
	FullName string
	PathToSignature string
}

//Слайс со списом подписывающихся
var signatories []*Signatory

func NewSignatory(
	serial	 		int,
	jobTitle 		string,
	fullName 		string,
	pathToSignature string,
) (*Signatory, error) {
	return &Signatory{serial, jobTitle, fullName, pathToSignature}, nil
}

//Формируем список подписывающихся
func AddSignatories(list  []string) ([]*Signatory, error) {
	for _, file := range list {
		fileSplit := strings.Split(file, "|")
		serialNumber, err :=  strconv.Atoi(fileSplit[0])
		if err != nil {
			log.Fatal(err)
		}
		signatory, err := NewSignatory(serialNumber, fileSplit[1],  fileSplit[2], file)
		if err != nil {
			log.Fatal(err)
		}

		signatories = append(signatories, signatory)
	}
	return signatories, nil
}
