package htmlgenerator

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/adminsemy/generatorReferanceList/interactive"
	"github.com/adminsemy/generatorReferanceList/signatories"
)


func Generate(path string, s []*signatories.Signatory, i *interactive.Answers) *strings.Reader{
	t, err := template.ParseFiles(path)
	if err != nil  {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, i); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
	return strings.NewReader(buf.String())
}