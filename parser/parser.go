package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func processCsv(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	i := 0
	for {
		row, error := r.Read()
		if error == io.EOF {
			break
		}
		if i == 1 {
			fmt.Println(row[6])
			break
		}
		i++
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("You have not entered the file path")
		return
	}
	fileName := os.Args[1]
	processCsv(fileName)

}
