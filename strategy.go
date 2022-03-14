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
