package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebServer struct {
	Router        *mux.Router
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        mux.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddRoute(method string, path string, handler http.HandlerFunc) {
	if s.Handlers[path] == nil {
		s.Handlers[path] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[path][method] = handler
	s.Router.HandleFunc(path, handler).Methods(method)
}

func (s *WebServer) Start() {
	http.ListenAndServe(":"+s.WebServerPort, s.Router)
}
