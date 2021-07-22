package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("error in lister")
		os.Exit(1)
	}

	defer li.Close()

	for {

		conn, err := li.Accept()
		if err != nil {
			log.Fatal("error in Accept")
			os.Exit(2)

		}
		//fmt.Printf("%T", li)
		//fmt.Printf("%T", conn)
		io.WriteString(conn, "hyyjjy")
		fmt.Fprintln(conn, "this is regarding Println")
		conn.Close()
	}

}
