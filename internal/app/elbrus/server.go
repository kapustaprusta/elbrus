package elbrus

import (
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

type track struct {
	name    string
	addedBy string
	date    string
	likes   int
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
}

func (s *server) handleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mainPageRawTmpl, err := ioutil.ReadFile("/home/vniksihov/src/golang/elbrus/web/template/index.html")
		if err != nil {
			s.logger.Errorln(err)

			return
		}

		mainPageTmpl := string(mainPageRawTmpl)
		parsedMainPageTmpl := template.Must(template.New("tmpl").Parse(mainPageTmpl))
		parsedMainPageTmpl.Execute(w, []track{})
	}
}
