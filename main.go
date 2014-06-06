package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		for i := 0; i < 18; i++ {
			err := PiGlow(byte(i), byte(rand.Int()))
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}
