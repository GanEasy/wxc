package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	/***以下是兼容前端的**/
	e.File("/", "dist/index.html")
	e.File("/live", "dist/index.html")
	e.File("/admin", "dist/index.html")
	e.Static("static", "dist/static")

	e.Logger.Fatal(e.Start(":9001"))
}
