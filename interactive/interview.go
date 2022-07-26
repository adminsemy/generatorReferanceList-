package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Структура для хранения ответов от пользователей
type Answers struct {
	FamiliarizationWith string //"Ознакомлен с" - здесь название приказа/распоряжения
	Date 				string //Дата ознакомления. В любом формате, не изменяется
	IdSignatories 		[]int //Список номеров (номера в названии подписи). Разделяем запятыми. Если 0 - то всех
}

func NewAnswers(signatories []string) (*Answers, error) {
	var answer Answers
	createFamiliarizationWith(&answer)
	createDate(&answer)
	return &answer, nil
}

func createFamiliarizationWith(answer *Answers) {
	fmt.Println("Введите полную строку, с чем нужно ознакомиться./n Например: с приказом №1 от 26.07.2022")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer.FamiliarizationWith = scanner.Text()
}

func createDate(answer *Answers) {
	fmt.Println("Введите дату ознакомления в нужном формате./n Например: 26.07.2022")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer.Date = scanner.Text()
}


func createIdSignatories(answer *Answers, signatories []string) {
	fmt.Println("Выберите людей для ознакомления. Необходимо вводить только номера./nЧисло 0 означает выбрать всех")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data, err := strconv.Atoi(scanner.Text())
	if  err != nil{
		data = 0
	}
	fmt.Println(data)
}