package main

import (
	"./handler"
	"./router"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"net/http"
)

func main() {

	commonHandlers := alice.New(context.ClearHandler, handler.LoggingHandler, handler.RecoverHandler)

	router := router.NewRouter()

	router.Get("/bowwo", commonHandlers.ThenFunc(handler.BowwoHandler))

	router.Get("/", commonHandlers.ThenFunc(handler.IndexHandler))

	http.ListenAndServe(":8000", router)
}
