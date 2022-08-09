package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	createIdSignatories(&answer, signatories)
	return &answer, nil
}

func createFamiliarizationWith(answer *Answers) {
	fmt.Printf("Введите полную строку, с чем нужно ознакомиться.\n Например: с приказом №1 от 26.07.2022\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer.FamiliarizationWith = scanner.Text()
}

func createDate(answer *Answers) {
	fmt.Printf("Введите дату ознакомления в нужном формате. Например: 26.07.2022 \nЕсли ничего не ввести поставит текущую дату\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer.Date = scanner.Text()
	if (answer.Date == "") {
		dt := time.Now()
		answer.Date = dt.Format("02.01.2006")
	}
}


func createIdSignatories(answer *Answers, signatories []string) {
	fmt.Printf("Выберите людей для ознакомления. Необходимо вводить только номера через запятую.\nЧисло 0 означает выбрать всех\n")
	for _, signatory := range signatories {
		fmt.Println(signatory)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choiceSignatories := strings.Split(scanner.Text(), ",")
	for _, chchoiceSignatory := range choiceSignatories {
		intSignatory, err := strconv.Atoi(chchoiceSignatory)
		if err != nil {
			answer.IdSignatories = []int{0}
		}
		answer.IdSignatories = append(answer.IdSignatories, intSignatory)
	}
}