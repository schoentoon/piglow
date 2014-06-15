package main

import (
	"testing"
	"time"
)

func TestSingleLed(t *testing.T) {
	PiGlowLed(1, Red, 0x10)
	ShutDown()
}

func TestCircleLegs(t *testing.T) {
	PiGlowLeg(0, 0x10)
	time.Sleep(500 * time.Millisecond)
	PiGlowLeg(0, 0x00)
	PiGlowLeg(1, 0x10)
	time.Sleep(500 * time.Millisecond)
	PiGlowLeg(1, 0x00)
	PiGlowLeg(2, 0x10)
	time.Sleep(500 * time.Millisecond)
	ShutDown()
}

func TestFlashRings(t *testing.T) {
	for i := 0; i < 10; i++ {
		for ring := 0; ring < 6; ring++ {
			if ring == 0 {
				PiGlowRing(5, 0x00)
			} else {
				PiGlowRing(byte(ring-1), 0x00)
			}
			PiGlowRing(byte(ring), 0x10)
			time.Sleep(100 * time.Millisecond)
		}
	}
	ShutDown()
}
