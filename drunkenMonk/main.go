package main

import (
	"fmt"
	"net/http"

	"strings"

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

type WineCountry struct {
	Country string   `json:country`
	Data    []string `json:data`
}
type WineCost struct {
	Data deets.WinePriceList `json:Data`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/countries", getCountries)
	e.GET("/wine/:country", winesInACountry)
	e.GET("/wines", getAllWines)
	e.GET("/maxcost", maxPriceList)
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
	limit := 10
	offset := c.QueryParam("offset")
	wines := deets.GetAllWines(limit, offset)
	u := &Wines{
		Data: wines,
	}
	return c.JSON(http.StatusOK, u)
}

func winesInACountry(c echo.Context) error {
	limit := 10
	offset := c.QueryParam("offset")
	country := c.Param("country")
	countryTitle := strings.Title(country)
	wines := deets.GetWinesInCountry(countryTitle, limit, offset)
	u := &WineCountry{
		Country: countryTitle,
		Data:    wines,
	}
	return c.JSON(http.StatusOK, u)
}

func maxPriceList(c echo.Context) error {
	offset := c.QueryParam("offset")
	wines := deets.Getmostexpensive(offset)
	u := &WineCost{
		Data: wines,
	}
	return c.JSON(http.StatusOK, u)

}
