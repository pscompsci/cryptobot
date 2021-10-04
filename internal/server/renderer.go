package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func (s *server) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s\n", err.Error(), debug.Stack())
	log.Println(trace)
}

func (s *server) render(w http.ResponseWriter, r *http.Request, name string) {
	ts, ok := s.templates[name]
	if !ok {
		s.serverError(w, fmt.Errorf("the template '%s' does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, nil)
	if err != nil {
		s.serverError(w, err)
	}
	buf.WriteTo(w)
}
