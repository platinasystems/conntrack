package main

// Example usage

import (
	"fmt"

	"github.com/platinasystems/conntrack"
	"log"
	"time"
	"syscall"
)

func main() {
/*	cs, err := conntrack.Established()
	if err != nil {
		panic(fmt.Sprintf("Established: %s", err))
	}
	fmt.Printf("Established on start:\n")
	for _, cn := range cs {
		fmt.Printf(" - %s\n", cn)
	}
	fmt.Println("")

	c, err := conntrack.New()
	if err != nil {
		panic(err)
	}
	for range time.Tick(1 * time.Second) {
		fmt.Printf("Connections:\n")
		for _, cn := range c.Connections() {
			fmt.Printf(" - %s\n", cn)
		}
	}*/
	for {
		conns, err := conntrack.Connections(syscall.AF_INET)
		if err != nil {
			log.Fatalln("failed to fetch connections...ERROR:", err)
		}
		for _, c := range conns {
			fmt.Println(c)
		}
		<- time.After(time.Millisecond * 300)
	}
}
