package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"opatutorial/middleware/bundler"

	manager "opatutorial/middleware/configurationmanager"
	"opatutorial/models/dao"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	e.Static("/static", "static")
	// e.GET("/", hello)
	// e.GET("/bundle", bundleTest)
	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
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
	// testGzip()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in loading .env file")
		panic(err)
	}

	if err := manager.Instance.ConnectDB(os.Getenv("SQL_DIALECT"), os.Getenv("SQL_DSN")); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("static", 0700); err != nil {
		panic(err)
	}
	bundler.InitCFManager(manager.Instance)
	dao.InitCFManager(manager.Instance)
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for t := range ticker.C {
			log.Println("Ticker at:", t)
			bundler.CreateBundleFile()
		}
	}()

	initWebServer()
}

// func testGzip() {
// 	flag.Parse() // get the arguments from command line

// 	destinationfile := flag.Arg(0)
// 	sourcedir := flag.Arg(1)

// 	if err := tarball.CompressTarball(destinationfile, sourcedir); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Success")
// }

func hello(c echo.Context) (erro error) {
	services, err := dao.ServiceDAO.GetAllServiceWithServiceGroupNameAndURL()
	if err != nil {
		erro = echo.ErrInternalServerError
		return
	}
	return c.JSON(http.StatusOK, services)
}
