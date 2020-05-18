package main

import (
	"fmt"

	"golang.org/x/net/context"
)

func (s *Server) runComputation(ctx context.Context) error {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	s.Log(fmt.Sprintf("Sum is %v", sum))
	return nil
}
