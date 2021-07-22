package main

import (
	"log"
	"os"
	"text/template"
)

var tp *template.Template

func init() {
	tp = template.Must(template.ParseFiles("myhtml.gohtml", "myhtml2.gohtml"))
}

// Main function
func main() {
	err := tp.ExecuteTemplate(os.Stdout, "myhtml2.gohtml", 556)
	if err != nil {
		log.Fatal("execute template error occur")
	}

	err = tp.ExecuteTemplate(os.Stdout, "myhtml.gohtml", `my name is arpit`)
	if err != nil {
		log.Fatal("execute template error occur")
	}

}
