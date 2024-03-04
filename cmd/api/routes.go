package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// not found handler}

	// method not allowed

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	return router
}
