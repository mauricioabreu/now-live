package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type API struct {
	*echo.Echo
	store *Store
}

func NewAPI() *API {
	return &API{
		echo.New(),
		NewStore(),
	}
}

func main() {
	api := NewAPI()
	api.GET("/live/:resource", api.GetResource)
	api.Logger.Fatal(api.Start(":1323"))
}

func (a *API) GetResource(c echo.Context) error {
	resource := c.Param("resource")
	object, err := a.store.GetObject(resource)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(object))
}
