package image

import (
	"net/http"
	"fmt"
        "github.com/ant0ine/go-urlrouter"
)

const GET = true

var Routes = []urlrouter.Route{
	urlrouter.Route{
		PathExp: "/api/v1/image",
		Dest: Root,
	},
	urlrouter.Route{
		PathExp: "/api/v1/image/:imagekey",
		Dest: Image,
	},
}

func Root(w http.ResponseWriter, req *http.Request, params map[string]string) {
        fmt.Fprintf(w, "Api Image Root") 
}

func Image(w http.ResponseWriter, req *http.Request, params map[string]string) {
        fmt.Fprintf(w, "Api Image %s", params["imagekey"]) 
}


