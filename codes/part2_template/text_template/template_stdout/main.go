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

	err = tf.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Execute error occur")
	}

}
