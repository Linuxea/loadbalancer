package lb

import (
	"net/http"
	"time"
)

type LB struct {
	port               int
	servers            []server
	NextServerStrategy stragety
}

func (l *LB) ServerProxy(rw http.ResponseWriter, r *http.Request) {

	for {
		s := l.NextServerStrategy.Next(l.servers)

		if s != nil && s.Alive() {
			s.Server(rw, r)
			return
		}
		time.Sleep(time.Microsecond * 300)
	}

}

func (l *LB) RegisterServer(s server) {
	l.servers = append(l.servers, s)
}
