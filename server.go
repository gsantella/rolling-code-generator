package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"io"
	"math"
	"math/big"
	"net/http"
	"os"
	"rolling-code-generator/namesgenerator"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
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

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.HideBanner = true

	uuidServiceKey = uuid.New().String()

	// print custom banner
	printCustomBanner()

	// initial tick
	tick()

	// run background ticking task
	go bgTask()

	// routes 
	e.GET("/", Hello)

	e.GET("/test", func(c echo.Context) error {
		output := "uuid: " + uuidServiceKey + " rollingCode: " + rollingCode + " randomInt: " + strconv.FormatInt(secureRandomInt, 10)
		return c.String(http.StatusOK, output)
	})

	e.Static("/", "static")

	e.Logger.Fatal(e.Start(":1324"))

}

func Hello(c echo.Context) error {
	data := struct{
		UuidServiceKey string
		RollingCode string
		SecureRandomInt string
	}{
		UuidServiceKey: uuidServiceKey,
		RollingCode: rollingCode,
		SecureRandomInt: strconv.FormatInt(secureRandomInt, 10),
	}
	return c.Render(http.StatusOK, "hello", data)
}


func tick() {
	secureRandomInt, _ = randint64()
	rollingCode  = namesgenerator.GetRandomName(0)
}

func bgTask() {
	ticker := time.NewTicker(time.Duration(updateNumSeconds) * time.Second)

	for range ticker.C {
		tick()
	}
}

func randint64() (int64, error) {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return 0, err
	}
	return val.Int64(), nil
}

func printCustomBanner() {

	fmt.Println("")
	fmt.Println("██████╗  ██████╗ ██████╗ ")
	fmt.Println("██╔══██╗██╔════╝██╔════╝ ")
	fmt.Println("██████╔╝██║     ██║  ███╗")
	fmt.Println("██╔══██╗██║     ██║   ██║")
	fmt.Println("██║  ██║╚██████╗╚██████╔╝")
	fmt.Println("╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ")
	fmt.Println("rolling-code-generator")
	fmt.Println("⇨ version 0.1")
	fmt.Println("⇨ service uuid " + uuidServiceKey)
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
