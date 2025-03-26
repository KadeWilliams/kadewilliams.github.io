package main

import (
	"io"
	"net/http"
	"text/template"

	go_templates "my-portfolio/go_templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Renderer = echo.New().Renderer
	e.Use(middleware.Logger())
	g := e.Group("")
	g.Use(middleware.BasicAuth(func(userName, password string, c echo.Context) (bool, error) {
		if userName == "kade" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	component := go_templates.Hello("John")
	e.GET("/hello", func(c echo.Context) error {
		return render(c, component)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}