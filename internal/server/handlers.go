package server

import "net/http"

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.render(w, r, "home.page.tmpl")
	}
}
