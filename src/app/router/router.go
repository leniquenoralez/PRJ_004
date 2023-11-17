package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Mux *mux.Router
}

func (r *Router) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.Mux.HandleFunc(path, f).Methods("GET")
}

func (r *Router) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.Mux.HandleFunc(path, f).Methods("POST")
}

func (r *Router) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.Mux.HandleFunc(path, f).Methods("PUT")
}

func (r *Router) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.Mux.HandleFunc(path, f).Methods("DELETE")
}

func (r *Router) Run(host string) {
	log.Fatal(http.ListenAndServe(host, r.Mux))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (r *Router) HandleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
func (r *Router) InitRouter() {
	r.Mux = mux.NewRouter()
}
