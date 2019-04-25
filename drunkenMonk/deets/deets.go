package deets

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllCountries() []string {
	var country string
	countryList := make([]string, 0)
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
		fmt.Println(country)
		countryList := append(countryList, country)
		fmt.Println(countryList)
	}
	return countryList
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
