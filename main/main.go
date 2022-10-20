package main

import (
	"circuit-breaker-go-example/client"
	"circuit-breaker-go-example/server"
	"fmt"
	"github.com/sony/gobreaker"
	"time"
)

func main() {
	go server.Server()

	// call with circuit breaker
	cb := gobreaker.NewCircuitBreaker(
		gobreaker.Settings{
			Name:        "my-circuit-breaker",
			MaxRequests: 3,
			Timeout:     3 * time.Second,
			Interval:    1 * time.Second,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return counts.ConsecutiveFailures >= 3 && failureRatio > 0.6
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				fmt.Printf("CircuitBreaker '%s' changed from '%s' to '%s'\n", name, from, to)
			},
		},
	)
	fmt.Println("Call with circuit breaker")
	for i := 0; i < 100; i++ {
		_, err := cb.Execute(func() (interface{}, error) {
			err := client.DoReq()
			return nil, err
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)

	}
}
