package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tf, err := template.ParseFiles("myhtml.gohtml")
	if err != nil {
		log.Fatal("error occur")
	}

	fp, err := os.Create("index.html")
	if err != nil {
		log.Fatal("file creation error")
	}
	defer fp.cedclose()
	err = tf.Execute(fp, nil)
	if err != nil {
		log.Fatal("Execute error occur")
	}

}
