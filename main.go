package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	template *template.Template
}

func (t *Templates) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		template: template.Must(template.ParseGlob("views/*.html")),
	}
}

func GetHome(c echo.Context) error {
	return c.Render(http.StatusOK, "index", struct{}{})
}

func main() {
	fmt.Println("starting")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("could not load .env file")
	}
	port := os.Getenv("PORT")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()

	e.GET("/", GetHome)

	e.Logger.Fatal(e.Start(":" + port))
}
