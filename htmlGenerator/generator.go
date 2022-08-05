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

type Data struct {
	Signatories []*signatories.Signatory
	Answers *interactive.Answers
	Paths *paths.Paths
}


func Generate(d *Data) *strings.Reader{
	t, err := template.New("referanceList.html").Funcs(template.FuncMap{
		"equal" : func (s *signatories.Signatory, a int)  bool{
			if (a == 0) {
				return true
			}
			return s.SerialNumber == a
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