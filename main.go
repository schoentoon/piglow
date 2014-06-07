package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 18; i++ {
		err := PiGlow(byte(i), byte(rand.Int()))
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
	ShutDown()
}
