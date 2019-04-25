package sqlCaller

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertIntoSql(country string, description string, designation string, price string, title string, winery string, taster_name string, taster_twitter_handle string) {
	db, err := sql.Open("mysql", "root:password@/winereviews?multiStatements=true")
	defer db.Close()
	checkError(err)
	stmt, err := db.Prepare("INSERT wineDetails SET country=?,details=?,designation=?,price=?,title=?,winery=?,taster_name=?")
	defer stmt.Close()
	checkError(err)
	res, err := stmt.Exec(country, description, designation, price, title, winery, taster_name)
	checkError(err)
	affect, err := res.RowsAffected()
	fmt.Println(affect)
	checkError(err)
}

func insertReviewer(email string) {

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
