package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"ic/valid"
)

func setupRouter() *echo.Echo {
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
		return c.Render(http.StatusOK, "form.html", nil)
	})

	e.POST("/user", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}

		if valid.IsEmpty(u.Email) {
			return c.JSON(http.StatusConflict, u)
		}

		if !valid.IsEmailValid(u.Email) {
			return c.JSON(http.StatusConflict, u)
		}

		return c.JSON(http.StatusCreated, u)
	})

	return e
}

func TestHomeHandler(t *testing.T) {
	e := setupRouter()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, e.Router().Find(http.MethodGet, "/", c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Hello, World!")
	}
}

func TestAccountHandler(t *testing.T) {
	e := setupRouter()
	req := httptest.NewRequest(http.MethodGet, "/account", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, e.Router().Find(http.MethodGet, "/account", c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUserHandler(t *testing.T) {
	e := setupRouter()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, e.Router().Find(http.MethodGet, "/user", c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Email validation")
	}
}

func TestUserPostHandler_ValidEmail(t *testing.T) {
	e := setupRouter()
	form := "email=test@example.com"
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(form))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, e.Router().Find(http.MethodPost, "/user", c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUserPostHandler_InvalidEmail(t *testing.T) {
	e := setupRouter()
	form := "email=invalid-email"
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(form))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, e.Router().Find(http.MethodPost, "/user", c)) {
		assert.Equal(t, http.StatusConflict, rec.Code)
	}
}

func TestUserPostHandler_EmptyEmail(t *testing.T) {
	e := setupRouter()
	form := "email="
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(form))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, e.Router().Find(http.MethodPost, "/user", c)) {
		assert.Equal(t, http.StatusConflict, rec.Code)
	}
}
