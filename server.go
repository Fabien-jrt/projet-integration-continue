package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// TemplateRenderer ...
type TemplateRenderer struct {
	templates *template.Template
}

// Render ...
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Set up template renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	e.GET("/account", func(c echo.Context) error {
		u := new(User)
		return c.JSON(http.StatusCreated, u)
	})

	e.GET("/user", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1>Create a new user!</h1>
		<form>
		
		</form>`)
	})

	e.POST("/user", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
