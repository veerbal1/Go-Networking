package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	str := `Hello Veerbal`

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
