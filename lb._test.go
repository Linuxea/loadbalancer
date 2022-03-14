package lb

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

type SimpleServer struct {
	message string
}

func (s *SimpleServer) Server(rw http.ResponseWriter, h *http.Request) {
	rw.Write([]byte(s.message))
}

func (*SimpleServer) Alive() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(100)
	return num > 30
}

var hello = &SimpleServer{message: "hello"}
var hi = &SimpleServer{message: "hi"}
var gb = &SimpleServer{message: "gb"}

var lb = &LB{
	port:               9090,
	NextServerStrategy: &SimplePoolStrategy{},
}

func init() {
	lb.RegisterServer(hello)
	lb.RegisterServer(hi)
	lb.RegisterServer(gb)
}

func TestLb(t *testing.T) {
	sm := http.NewServeMux()
	sm.HandleFunc("/", lb.ServerProxy)
	err := http.ListenAndServe(fmt.Sprintf(":%d", lb.port), sm)
	handlerError(err)
}
