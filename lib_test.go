/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <nighteyes1993@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return Toon Schoenmakers
 * ----------------------------------------------------------------------------
 */

package piglow

import (
	"testing"
	"time"
)

func TestSingleLed(t *testing.T) {
	if HasPiGlow() == false {
		t.Fatal("No piglow detected")
	}

	Led(1, Red, 0x10)
	ShutDown()
}

func TestCircleLegs(t *testing.T) {
	if HasPiGlow() == false {
		t.Fatal("No piglow detected")
	}

	Leg(0, 0x10)
	time.Sleep(500 * time.Millisecond)
	Leg(0, 0x00)
	Leg(1, 0x10)
	time.Sleep(500 * time.Millisecond)
	Leg(1, 0x00)
	Leg(2, 0x10)
	time.Sleep(500 * time.Millisecond)
	ShutDown()
}

func TestFlashRings(t *testing.T) {
	if HasPiGlow() == false {
		t.Fatal("No piglow detected")
	}

	for i := 0; i < 10; i++ {
		for ring := 0; ring < 6; ring++ {
			if ring == 0 {
				Ring(5, 0x00)
			} else {
				Ring(byte(ring-1), 0x00)
			}
			Ring(byte(ring), 0x10)
			time.Sleep(100 * time.Millisecond)
		}
	}
	ShutDown()
}

func TestFadeLed(t *testing.T) {
	if HasPiGlow() == false {
		t.Fatal("No piglow detected")
	}

	Fade(0, Red, 0x00, 0x64, 10*time.Millisecond)
	ShutDown()
}
