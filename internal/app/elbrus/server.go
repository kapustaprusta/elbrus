package elbrus

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	s.configure()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configure() {
	s.router.HandleFunc("/", s.handleMain()).Methods("GET")
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/home/vniksihov/src/golang/elbrus/web/static"))))
}

func (s *server) handleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mainPageRaw, err := ioutil.ReadFile("/home/vniksihov/src/golang/elbrus/web/template/index.html")
		if err != nil {
			s.logger.Errorln(err)

			return
		}

		w.Write(mainPageRaw)
	}
}
