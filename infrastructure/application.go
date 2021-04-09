package infrastructure

import (
	"fmt"
	"log"
	"net/http"
)

type Application struct {
	router *Router
	port   string
	db     *DB
}

func NewApplication(port string) *Application {
	db := GetDB()
	router := NewRouter()
	return &Application{
		router: router,
		db:     db,
		port:   port,
	}
}

func (a *Application) RegisterRoute(method, path string, handler func(http.ResponseWriter, *http.Request) error) {
	a.router.AddRoute(Route{
		Method:     method,
		Path:       path,
		HandleFunc: handler,
	})
}

func (a *Application) Run() {
	fmt.Println("Server started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		route := a.router.FindRoute(r.Method, r.URL.Path)
		route.HandleFunc(w, r)
	})
	if err := http.ListenAndServe(a.port, nil); err != nil {
		log.Fatal(err)
	}
}
