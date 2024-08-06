package main

import (
	"html/template"
	"io"
	"net/http"
	"rolling-code-generator/namesgenerator"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	updateNumSeconds int = 5
	secureRandomInt int64
	rollingCode string
	uuidServiceKey string
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {

	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	t := &Template{
		templates: template.Must(template.ParseGlob("./public/views/*.html")),
		//templates: template.Must(template.ParseGlob("/usr/local/bin/public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.HideBanner = true

	// generate new UUID for service key
	uuidServiceKey = uuid.New().String()

	// print custom banner
	printCustomBanner()

	// initializing work
	tick()

	// run continual work at specified interval
	go tickRunner()

	// routes 
	e.GET("/", homeHandler)

	e.GET("/api", apiHandler)

	e.Static("/", "static")

	// commence logging
	logger.Info("App starting")

	e.Logger.Info(e.Start(":1324"))

	// todo: graceful shutdown, but probably not needed

}

func homeHandler(c echo.Context) error {
	data := struct{
		UuidServiceKey string
		RollingCode string
		SecureRandomInt string
	}{
		UuidServiceKey: uuidServiceKey,
		RollingCode: rollingCode,
		SecureRandomInt: strconv.FormatInt(secureRandomInt, 10),
	}
	return c.Render(http.StatusOK, "home", data)
}

func apiHandler(c echo.Context) error {
	data := struct{
		UuidServiceKey string
		RollingCode string
		SecureRandomInt string
	}{
		UuidServiceKey: uuidServiceKey,
		RollingCode: rollingCode,
		SecureRandomInt: strconv.FormatInt(secureRandomInt, 10),
	}
	return c.JSON(http.StatusOK, data)
}

func tick() {
	secureRandomInt, _ = getSecureRandInt64()
	rollingCode  = namesgenerator.GetRandomName(0)
}

func tickRunner() {
	ticker := time.NewTicker(time.Duration(updateNumSeconds) * time.Second)

	for range ticker.C {
		tick()
	}
}
