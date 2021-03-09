package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]

	htmlString := fmt.Sprintf(`
	 <!DOCTYPE html>
	 <html lang = "en">
	 <head>
	 <meta charset = "UTF-8">
	 <title> Hello World </title>
	 </head>
	 <body>

	 <h1>` + name + ` </h1>
	 </body>
	 </html>  	
	`)

	//fmt.Println(htmlString)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("some error occur")
		os.Exit(1)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(htmlString))

}
