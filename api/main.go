package main

import (
	"mime"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func init() {
	mime.AddExtensionType(".ts", "video/mp2t")
	mime.AddExtensionType(".m3u8", "application/vnd.apple.mpegurl")
	mime.AddExtensionType(".mp4", "video/mp4")
}

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
	api.GET("/live/*", api.GetResource)
	api.Logger.Fatal(api.Start(":1323"))
}

func (a *API) GetResource(c echo.Context) error {
	resource := Resource{Path: c.Request().RequestURI}
	object, err := a.store.GetObject(resource.ObjectName())
	if len(object) == 0 {
		c.Response().WriteHeader(http.StatusNotFound)
		return err
	}

	if err != nil {
		return err
	}

	contentType := mime.TypeByExtension(filepath.Ext(resource.Path))
	c.Response().Header().Set(echo.HeaderContentType, contentType)
	return c.String(http.StatusOK, string(object))
}
