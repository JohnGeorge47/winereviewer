package deets

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Wine struct {
	name string
	cost string
}
type WinePriceList struct {
	Wines []Wine
}

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

func GetWinesInCountry(country string, limit int, offset string) []string {
	if offset == "" {
		offset = "0"
	}
	wineList := make([]string, 0)
	var title string
	db, err := sql.Open("mysql", "root:password@/winereviews?multiStatements=true")
	defer db.Close()
	checkError(err)
	stmt, err := db.Prepare("SELECT title from wineDetails WHERE country=? LIMIT ? OFFSET ?")
	checkError(err)
	defer stmt.Close()
	rows, err := stmt.Query(country, limit, offset)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&title)
		wineList = append(wineList, title)
		fmt.Println(wineList)
	}
	return wineList
}

func Getmostexpensive(offset string) WinePriceList {
	if offset == "" {
		offset = "0"
	}
	var winess WinePriceList
	winess.Wines = make([]Wine, 0)
	var title string
	var price string
	db, err := sql.Open("mysql", "root:password@/winereviews?multiStatements=true")
	defer db.Close()
	checkError(err)
	stmt, err := db.Prepare("SELECT title,price FROM wineDetails OFFSET ? ORDER BY DESC")
	checkError(err)
	defer stmt.Close()
	rows, err := stmt.Query(offset)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&title, &price)
		wineDeets := Wine{name: title, cost: price}
		winess.Wines = append(winess.Wines, wineDeets)
	}
	fmt.Println(winess)
	return winess
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
