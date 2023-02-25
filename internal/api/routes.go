package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//  create url
//  go to url
//  delete url

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.Main)
	router.HandlerFunc(http.MethodGet, "/mini/gotourl", app.GoToURL)
	router.HandlerFunc(http.MethodPost, "/mini/create", app.CreateURL)
	router.HandlerFunc(http.MethodDelete, "/mini/delete/:id", app.DeleteURL)

	return router
}
