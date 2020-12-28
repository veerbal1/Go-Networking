package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	str := "Hello " + os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	fl, err := os.Create("index.txt")
	if err != nil {
		log.Fatal(err)
	}
	wr, error := io.Copy(fl, strings.NewReader(str))
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(wr)
}
