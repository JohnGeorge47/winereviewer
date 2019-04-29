package main

import (
	"fmt"
	"net/http"

	"github.com/JohnGeorge47/winereviewer/drunkenMonk/deets"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}
type Countries struct {
	Data []string `json:data`
}

type Wines struct {
	Data []string `json:data`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/countries", getCountries)
	e.GET("/wines", getAllWines)
	e.Logger.Fatal(e.Start(":1323"))
}

func getCountries(c echo.Context) error {
	countries := deets.GetAllCountries()
	fmt.Println(countries)
	u := &Countries{
		Data: countries,
	}
	return c.JSON(http.StatusOK, u)
}

func getAllWines(c echo.Context) error {
	wines := deets.GetAllWines()
	u := &Wines{
		Data: wines,
	}
	return c.JSON(http.StatusOK, u)
}
