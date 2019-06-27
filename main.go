package main

import (
	"flag"
	"fmt"
	"net/http"

	"opatutorial/models"
	"opatutorial/models/dao"
	"opatutorial/utils/tarball"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := models.ConnectDatabase(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// if err := models.MigrateDB(); err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	fmt.Println(dao.UserDAO.GetUserFromID())

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func testGzip() {
	flag.Parse() // get the arguments from command line

	destinationfile := flag.Arg(0)
	sourcedir := flag.Arg(1)

	if err := tarball.Compress_tarball(destinationfile, sourcedir); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, models.User{
		ID:   1,
		Name: "Sang LX",
	})
}
