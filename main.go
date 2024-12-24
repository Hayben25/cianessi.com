package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate(name string, file_location string) *template.Template {
	temp, err := template.New(name).ParseFiles(file_location)

	if err != nil {
		panic(err)
	}

	return temp
}

func Navbar(c echo.Context) error {
	return c.Render(http.StatusOK, "navbar", nil)
}
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}
func NasLand(c echo.Context) error {
	return c.Render(http.StatusOK, "nas-landing", nil)
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("src/views/*.html")),
	}
	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.File("/output.css", "src/output.css")
	e.File("/htmx.min.js", "htmx.min.js")
	e.GET("/nas-landing", NasLand)
	e.GET("/", Home)
	e.Logger.Fatal(e.Start(":42069"))
}
