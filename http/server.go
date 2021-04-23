package http

import (
	"net/http"

	"github.com/Rodrigo001-de/iniciando-golang/model"
	"github.com/labstack/echo"
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	// criando um novo Web Server
	e := echo.New()
	// quando alguém acessar /product vai rodar uma função chamada getAll
	e.GET("/product", w.getAll)
	e.POST("/product", w.createProduct)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(c echo.Context) error {
	// vai retornar os produtos em formato JSON
	return c.JSON(http.StatusOK, w.Products)
}

func (w WebServer) createProduct(c echo.Context) error {
	product := model.NewProduct()

	// o c.Bind é pegar os dados que eu esotu recebendo a request e grudar com a
	//  struct
	if err := c.Bind(product); err != nil {
		return err
	}
	w.Products.Add(*product)
	return c.JSON(http.StatusCreated, product)

}
