package server

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (s *server) routes() http.Handler {
	router := pat.New()

	router.Get("/", s.handleHome())

	fileserver := http.FileServer(http.Dir("./web/static"))
	router.Get("/static/", http.StripPrefix("/static", fileserver))

	return router
}
