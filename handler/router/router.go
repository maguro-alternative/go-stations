package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"

	"github.com/justinas/alice"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// create a *service.TODOService type variable using the *sql.DB type variable
	var todoService = service.NewTODOService(todoDB)

	// register routes
	mux := http.NewServeMux()
	middleChain := alice.New(
		middleware.Recovery,
		middleware.GetOSHandler,
		middleware.GetLog,
	)
	mux.Handle("/healthz", middleChain.Then(handler.NewHealthzHandler()))
	mux.Handle("/todos", middleChain.Then(handler.NewTODOHandler(todoService)))
	mux.Handle("/do-panic", middleChain.Then(handler.NewDoPanicHandler()))
	return mux
}
