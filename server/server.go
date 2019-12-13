package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server : server structure representin Http Server
type Server struct {
	ListenHost string
	ListenPort string
	Router     *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) handleWrongInput(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wrong data input, redirect to main page after 5 sec"))
}

func (s *Server) handleRootGetMethod(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index/index.html")
}

func (s *Server) handleRootPostMethod(w http.ResponseWriter, r *http.Request) {
	name, address := r.FormValue("name"), r.FormValue("address")
	token := name + ":" + address
	setCookie := http.Cookie{
		Name:  "token",
		Value: token,
	}
	http.SetCookie(w, &setCookie)
	http.ServeFile(w, r, "index/index.html")
}

// NewHTTPServer : create new Server on IP and Port
func NewHTTPServer(listenIP, httpPort string) *Server {
	res := Server{
		ListenHost: listenIP,
		ListenPort: httpPort,
		Router:     mux.NewRouter(),
	}
	res.Router.HandleFunc("/", res.handleRootGetMethod).Methods(http.MethodGet)
	res.Router.HandleFunc("/", res.handleRootPostMethod).Methods(http.MethodPost)
	return &res
}

// Serve : starts listening on IP and Port
func (s *Server) Serve() {
	listenString := s.ListenHost + ":" + s.ListenPort
	log.Println("Serving http://" + listenString)
	log.Fatal(http.ListenAndServe(listenString, s.Router))
}
