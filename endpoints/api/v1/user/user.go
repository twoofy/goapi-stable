package user

import (
	"net/http"
	"fmt"
        "github.com/ant0ine/go-urlrouter"
)

const GET = true

var Routes = []urlrouter.Route{
	urlrouter.Route{
		PathExp: "/api/v1/user",
		Dest: Root,
	},
}

func Root(w http.ResponseWriter, req *http.Request, params map[string]string) {
        fmt.Fprintf(w, "Api User Root") 
}


