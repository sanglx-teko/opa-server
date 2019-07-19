package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sanglx-teko/opa-server/middleware/bundler"

	manager "github.com/sanglx-teko/opa-server/middleware/configurationmanager"
	"github.com/sanglx-teko/opa-server/models/dao"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
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

func hello(c echo.Context) (erro error) {
	services, err := dao.ServiceDAO.GetAllServiceWithServiceGroupNameAndURL()
	if err != nil {
		erro = echo.ErrInternalServerError
		return
	}
	return c.JSON(http.StatusOK, services)
}
