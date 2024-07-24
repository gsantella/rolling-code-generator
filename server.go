package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"io"
	"math"
	"math/big"
	"net/http"
	"rolling-code-generator/namesgenerator"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

// todo: customized update frequency
//var updateTimeSeconds int = 5
var randomInt int64
var rollingCode string
var key string

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	key = uuid.New().String()

	go bgTask()

	e.GET("/", Hello)

	e.GET("/test", func(c echo.Context) error {
		output := "uuid: " + key + " rollingCode: " + rollingCode + " randomInt: " + strconv.FormatInt(randomInt, 10)
		return c.String(http.StatusOK, output)
	})

	e.Static("/", "static")

	e.Logger.Fatal(e.Start(":1324"))
}

func Hello(c echo.Context) error {
	data := struct{
		Uuid string
		RollingCode string
		Key string
	}{
		Uuid: key,
		RollingCode: rollingCode,
		Key: strconv.FormatInt(randomInt, 10),
	}
	return c.Render(http.StatusOK, "hello", data)
}


func bgTask() {
	ticker := time.NewTicker(5 * time.Second)

	for t := range ticker.C {
		randomInt, _ = randint64()
		rollingCode  = namesgenerator.GetRandomName(0)
		fmt.Println("Congrats ", rollingCode, " " , randomInt, " @ ", t)
	}
}

func randint64() (int64, error) {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return 0, err
	}
	return val.Int64(), nil
}
