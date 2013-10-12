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
	"github.com/twoofy/goapi-stable/endpoints/api"
	"github.com/twoofy/goapi-stable/endpoints/api/v1"
	"github.com/twoofy/goapi-stable/endpoints/api/v1/user"
	"github.com/twoofy/goapi-stable/endpoints/api/v1/image"
)

const POST string = "POST"
const GET string = "GET"
const PUT string = "PUT"
const OPTIONS string = "OPTIONS"
const DELETE string = "DELETE"

var runPort = flag.String("port", "", "port to run server on")

func main() {
	flag.Parse()

	routes := []urlrouter.Route{}
	routes = append(routes, api.Routes...)
	routes = append(routes, v1.Routes...)
	routes = append(routes, user.Routes...)
	routes = append(routes, image.Routes...)

	router := urlrouter.Router{
		Routes: routes,
	}

	router.Start()


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == GET {
			// placeholder
		}
		fmt.Printf("process %s request\n", r.Method)
		route, params := router.FindRouteFromURL(r.URL)
		handler := route.Dest.(func(http.ResponseWriter, *http.Request, map[string]string))
		handler(w, r, params)
	})
	fmt.Printf("running %s\n", *runPort)

	log.Fatal(http.ListenAndServe(":" + *runPort, nil))
}
