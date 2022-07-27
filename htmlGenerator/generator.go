package htmlgenerator

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"
)


func Generate(path string) *strings.Reader{
	t, err := template.ParseFiles(path)
	if err != nil  {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
	return strings.NewReader(buf.String())
}