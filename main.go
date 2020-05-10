package main

import (
	"os"
	"strings"

	"github.com/pignuante/test-crawler/scrapper"

	"github.com/labstack/echo"
	"github.com/pignuante/test-crawler/utils"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("./index.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(utils.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1190"))
	// scrapper.Scrape("term")
}
