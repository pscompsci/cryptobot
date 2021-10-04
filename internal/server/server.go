package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pscompsci/cryptobot/pkg/events"
)

type server struct {
	eventBus  *events.EventBus
	templates map[string]*template.Template
}

func New(eb *events.EventBus) *server {
	ts, err := newTemlateCache("./web/templates/")
	if err != nil {
		log.Fatal("could not parse templates")
		os.Exit(1)
	}

	return &server{
		eventBus:  eb,
		templates: ts,
	}
}

func (s *server) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", 5050),
		Handler:      s.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return srv.ListenAndServe()
}
