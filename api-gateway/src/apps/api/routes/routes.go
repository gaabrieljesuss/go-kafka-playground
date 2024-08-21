package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router interface {
	Load(*echo.Group)
}

type router struct {
}

func New() Router {
	return &router{}
}

func (r *router) Load(rootEndpoint *echo.Group) {
	r.LoadDocs(rootEndpoint)

	NewAccountRouter().Load(rootEndpoint)
}

func (r *router) LoadDocs(group *echo.Group) {
	group.GET("/docs/*", echoSwagger.WrapHandler)
	group.GET("/docs", func(c echo.Context) error {
		return c.Redirect(301, "/api/docs/index.html")
	})
}
