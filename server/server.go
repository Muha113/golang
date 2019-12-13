package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server : struct
type Server struct {
	ListenHost string
	ListenPort string
	Router     *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) myHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Host       string      `json:"host"`
		UserAgent  string      `json:"user_agent"`
		RequestURI string      `json:"request_uri"`
		Headers    http.Header `json:"headers"`
	}

	newResponse := response{
		r.Host,
		r.UserAgent(),
		r.RequestURI,
		r.Header,
	}
	js, err := json.Marshal(newResponse)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// NewHTTPServer : create new Server on IP and Port
func NewHTTPServer(listenIP, httpPort string) *Server {
	res := Server{
		ListenHost: listenIP,
		ListenPort: httpPort,
		Router:     mux.NewRouter(),
	}
	res.Router.HandleFunc("/drive", res.myHandler).Methods(http.MethodGet)
	return &res
}

// Serve : starts listening on IP and Port
func (s *Server) Serve() {
	listenString := s.ListenHost + ":" + s.ListenPort
	log.Println("Serving http://" + listenString)
	log.Fatal(http.ListenAndServe(listenString, s.Router))
}
