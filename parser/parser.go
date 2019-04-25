package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/JohnGeorge47/winereviewer/parser/sqlCaller"
)

func processCsv(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	for {
		row, error := r.Read()
		if error == io.EOF {
			break
		}
		fmt.Println(row)
		sqlCaller.InsertIntoSql(row[1], row[2], row[3], row[5], row[11], row[13], row[9], row[10])
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
