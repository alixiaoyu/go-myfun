package main

import (
	"fmt"
)

type Node struct {
	id     string
	ipaddr string
}

type Server struct {
	nodes [5]*Node
	num   int
}

func (s *Server) Init(n int, nodes []Node) {
	s.num = n
	for idx, v := range nodes {
		fmt.Println("====>", &idx, &v)
		s.nodes[idx] = &v
		// s.nodes[idx] = &nodes[idx]
	}
}

func (s *Server) Print() {
	for idx, v := range s.nodes {
		if v != nil {
			fmt.Printf("%d) %s--%s\n", idx, v.id, v.ipaddr)
		}
	}
}

func main() {
	nodes := []Node{
		Node{"0", "192.168.0.0.0"},
		Node{"1", "192.168.0.0.1"},
		Node{"2", "192.168.0.0.2"},
	}

	var s Server
	s.Init(3, nodes)
	s.Print()
}
