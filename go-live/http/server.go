package http

import (
	"github.com/labstack/echo/v4"
	"github.com/rscampos3/GoLang/tree/main/go-live/model"
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Server() {
	e := echo.New()
	e.GET("/product", w.getAll)
}

func (e WebServer) getAll(c echo.Context) error {

}
