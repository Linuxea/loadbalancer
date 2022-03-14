lb.go
```golang
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
```



server.go
```golang
package lb

import "net/http"

type server interface {
	Server(http.ResponseWriter, *http.Request)

	Alive() bool
}
```


strategy.go
```golang
package lb

type stragety interface {
	Next(servers []server) server
}

type SimplePoolStrategy struct {
	rountCount int
}

func (s *SimplePoolStrategy) Next(servers []server) server {

	if len(servers) == 0 {
		log("no server find")
		return nil
	}

	s.rountCount = s.rountCount + 1
	return servers[int(s.rountCount%len(servers))]
}
```


log.go
```golang
package lb

import (
	"fmt"
	"time"
)

func log(message string) {
	fmt.Println(fmt.Sprintf("%s - %s", time.Now().String(), message))
}
```


err_handler.go
```golang
package lb

import (
	"os"
)

func handlerError(err error) {

	log(err.Error())
	os.Exit(1)
}
```


lb_test.go
```golang
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
```


## 参考

- [1] [Building a Load Balancer in Go](https://medium.com/better-programming/building-a-load-balancer-in-go-3da3c7c46f30)
- [2] [https://github.com/Linuxea/loadbalancer](https://github.com/Linuxea/loadbalancer)