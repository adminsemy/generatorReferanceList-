package htmlgenerator

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/adminsemy/generatorReferanceList/interactive"
	"github.com/adminsemy/generatorReferanceList/paths"
	"github.com/adminsemy/generatorReferanceList/signatories"
)

//Сохраняем все данные для генерации листа ознакомления
type Data struct {
	Signatories []*signatories.Signatory
	Answers *interactive.Answers
	Paths *paths.Paths
}

//Генерируем лист для сохранения его в нужном формате
func Generate(d *Data) *strings.Reader{
	//Формируем счетчик для шаблона листа ознакомления
	var i int
	l := &i
	t, err := template.New("referanceList.html").Funcs(template.FuncMap{
		//Функция для определения, нужена ли эта подпись в листе ознакомления
		"equal" : func (s *signatories.Signatory, a int)  bool{
			if (a == 0) {
				return true
			}
			return s.SerialNumber == a
		},
		//Увеличиваем счетчик листа ознакомления и возвращаем результат
		"increment" : func () int{
			*l++
			return i
		},
	}).ParseFiles(d.Paths.PathToTemplate)
	if err != nil  {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, d); err != nil {
		log.Fatal(err)
	}
	return strings.NewReader(buf.String())
}