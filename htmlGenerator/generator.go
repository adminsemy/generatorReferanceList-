package htmlgenerator

import (
	"bytes"
	"html/template"
	"log"
	"strings"

	"github.com/adminsemy/generatorReferanceList/interactive"
	"github.com/adminsemy/generatorReferanceList/signatories"
)

type Data struct {
	Signatories []*signatories.Signatory
	Answers *interactive.Answers
	CurrentDirectory string
}


func Generate(path string, d *Data) *strings.Reader{
	t, err := template.ParseFiles(path)
	if err != nil  {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, d); err != nil {
		log.Fatal(err)
	}
	return strings.NewReader(buf.String())
}