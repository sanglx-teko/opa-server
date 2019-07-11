package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"opatutorial/utils/tarball"

	"opatutorial/middleware/bundler"
	manager "opatutorial/middleware/configurationmanager"

	"opatutorial/models/dao"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initWebServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())

	// Routes
	e.GET("/", hello)
	e.GET("/bundle", bundleTest)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func bundleTest(c echo.Context) (erro error) {
	userRoles, err := dao.UserDAO.GetAllUserWithRoles()
	if err != nil {
		erro = echo.ErrInternalServerError
		return
	}
	return c.JSON(http.StatusOK, userRoles)
}

func main() {

	if err := manager.Instance.ConnectDB("mysql", "root:123@tcp(localhost:3306)/opadb?parseTime=true&charset=utf8"); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("static", 0700); err != nil {
		panic(err)
	}
	bundler.InitCFManager(manager.Instance)
	dao.InitCFManager(manager.Instance)
	bundler.CreateBundleFile()

	initWebServer()
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

func hello(c echo.Context) (erro error) {
	services, err := dao.ServiceDAO.GetAllServiceWithServiceGroupNameAndURL()
	if err != nil {
		erro = echo.ErrInternalServerError
		return
	}
	return c.JSON(http.StatusOK, services)
}
