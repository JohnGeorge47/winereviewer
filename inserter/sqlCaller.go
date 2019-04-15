package sqlCaller

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InsertIntoSql(country string, description string, designation string, title string, winery string, taster_name string, taster_twitter_handle string) error {
	db, err := sql.Open("mysql", "root:password@/winedata?multiStatements=true")
	defer db.Close()
	checkError(err)
	tx, err := db.Begin()
	checkError(err)
	{
		stmt, err := tx.Prepare("INSERT wineDetails SET country=?,description=?,designation=?,title=?,winery=?")
		if err != nil {
			tx.Rollback()
			return err
		}

		defer stmt.Close()
		if _, err := stmt.Exec(country, description, designation, title, winery); err != nil {
			tx.Rollback()
			return err
		}
	}
	{
		stmt, err := tx.Prepare("INSERT into reviewer SET taster_name=?,taster_twitter_handle=?")
		if err != nil {
			tx.Rollback()
			return err
		}
		if _, err := stmt.Exec(taster_name, taster_twitter_handle); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func insertReviewer(email string) {

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
