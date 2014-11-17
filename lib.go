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
	"bitbucket.org/gmcbay/i2c"
	"errors"
	"time"
)

var bus, bus_err = i2c.Bus(1)

const (
	address = 0x54

	enableOutput = 0x00
	enableLeds   = 0x13
	setPwmValues = 0x01
	update       = 0x16
)

// All the possible colors
const (
	Red byte = iota
	Orange
	Yellow
	Green
	Blue
	White
)

func init() {
	if bus != nil && bus_err == nil {
		bus_err = bus.WriteByte(address, enableOutput, 0x01)
		bus_err = bus.WriteByteBlock(address, enableLeds, []byte{0xFF, 0xFF, 0xFF})
	}
}

var values = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

// Array with the bytes for all the individual leds mapped according to [leg][color]
var legs = [][]byte{{6, 7, 8, 5, 4, 9}, {17, 16, 15, 13, 11, 10}, {0, 1, 2, 3, 14, 12}}

// PiGlow toggle a single led to a certain intensity
func PiGlow(led, intensity byte) error {
	if bus_err != nil {
		return bus_err
	}

	if led < 0 || led > 18 {
		return errors.New("invalid LED")
	}

	values[led] = intensity

	err := bus.WriteByteBlock(address, setPwmValues, values)
	if err != nil {
		return err
	}
	err = bus.WriteByte(address, update, 0xFF)
	return err
}

// Led toggle a single led based on a leg and color
func Led(leg, color, intensity byte) error {
	if bus_err != nil {
		return bus_err
	}

	if leg < 0 || leg > 2 {
		return errors.New("invalid leg")
	}
	if color < 0 || color > 5 {
		return errors.New("invalid color")
	}

	return PiGlow(legs[leg][color], intensity)
}

// Leg entirely light up a complete leg
func Leg(leg, intensity byte) error {
	if bus_err != nil {
		return bus_err
	}

	if leg < 0 || leg > 2 {
		return errors.New("invalid leg")
	}

	for _, led := range legs[leg] {
		values[led] = intensity
	}

	err := bus.WriteByteBlock(address, setPwmValues, values)
	if err != nil {
		return err
	}
	err = bus.WriteByte(address, update, 0xFF)
	return err
}

// Ring entirely light up a certain color/ring
func Ring(color, intensity byte) error {
	if bus_err != nil {
		return bus_err
	}

	if color < 0 || color > 5 {
		return errors.New("invalid ring")
	}

	values[legs[0][color]] = intensity
	values[legs[1][color]] = intensity
	values[legs[2][color]] = intensity

	err := bus.WriteByteBlock(address, setPwmValues, values)
	if err != nil {
		return err
	}
	err = bus.WriteByte(address, update, 0xFF)
	return err
}

// Fade a certain led at leg with color from intensity from to intensity to
// with intervals of interval
func Fade(leg, color, from, to byte, interval time.Duration) error {
	if bus_err != nil {
		return bus_err
	}

	if leg < 0 || leg > 2 {
		return errors.New("invalid leg")
	}
	if color < 0 || color > 5 {
		return errors.New("invalid color")
	}

	for i := from; i != to; i++ {
		values[legs[leg][color]] = i
		err := bus.WriteByteBlock(address, setPwmValues, values)
		if err != nil {
			return err
		}
		err = bus.WriteByte(address, update, 0xFF)
		if err != nil {
			return err
		}
		time.Sleep(interval)
	}
	return nil
}

// ShutDown Turn off all the lights
func ShutDown() error {
	if bus_err != nil {
		return bus_err
	}

	for i := 0; i < 18; i++ {
		values[i] = 0x00
	}

	err := bus.WriteByteBlock(address, setPwmValues, values)
	if err != nil {
		return err
	}
	err = bus.WriteByte(address, update, 0xFF)

	return err
}
