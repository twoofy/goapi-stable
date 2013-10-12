package api

import (
	"net/http"
	"fmt"
        "github.com/ant0ine/go-urlrouter"
)

const GET = true

var Routes = []urlrouter.Route{
	urlrouter.Route{
		PathExp: "/api",
		Dest: Root,
	},
}

func Root(w http.ResponseWriter, req *http.Request, params map[string]string) {
        fmt.Fprintf(w, "Api Root") 
}


