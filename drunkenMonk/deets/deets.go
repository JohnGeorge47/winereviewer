package deets

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllCountries() []string {
	var country string
	countryList := make([]string, 0, 100)
	db, err := sql.Open("mysql", "root:password@/winereviews?multiStatements=true")
	defer db.Close()
	checkError(err)
	stmt, err := db.Prepare("SELECT DISTINCT country from wineDetails")
	checkError(err)
	defer stmt.Close()
	rows, err := stmt.Query()
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&country)
		countryList = append(countryList, country)
		fmt.Println(countryList)
	}
	return countryList
}

func GetAllWines(limit int, offset string) []string {
	if offset == "" {
		offset = "0"
	}
	var title string
	wineList := make([]string, 0)
	// query := fmt.Sprintf("SELECT title from wineDetailscodecod")
	db, err := sql.Open("mysql", "root:password@/winereviews?multiStatements=true")
	defer db.Close()
	checkError(err)
	stmt, err := db.Prepare("SELECT title from wineDetails LIMIT ? OFFSET ?")
	checkError(err)
	defer stmt.Close()
	rows, err := stmt.Query(limit, offset)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&title)
		wineList = append(wineList, title)
		fmt.Println(wineList)
	}
	return wineList
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
