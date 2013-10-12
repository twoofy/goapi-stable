package main

import (
	"flag"
	"fmt"
	"github.com/ant0ine/go-urlrouter"
	"log"
	"net/http"
)

// Individual endpoints
import (
	"api/endpoints/api/v1"
	"api/endpoints/api/v1/user"
	"api/endpoints/api/v1/image"
)

var runPort = flag.String("port", "", "port to run server on")

func main() {
	flag.Parse()

	routes := []urlrouter.Route{}
	routes = append(routes, v1.Routes...)
	routes = append(routes, user.Routes...)
	routes = append(routes, image.Routes...)

	router := urlrouter.Router{
		Routes: routes,
	}

	router.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		route, params := router.FindRouteFromURL(r.URL)
		handler := route.Dest.(func(http.ResponseWriter, *http.Request, map[string]string))
		handler(w, r, params)
	})
	fmt.Printf("running %s\n", *runPort)

	log.Fatal(http.ListenAndServe(":" + *runPort, nil))
}
