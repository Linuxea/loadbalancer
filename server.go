package lb

import "net/http"

type server interface {
	Server(http.ResponseWriter, *http.Request)

	Alive() bool
}
