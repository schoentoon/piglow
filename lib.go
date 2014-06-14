package main

import (
	"bitbucket.org/gmcbay/i2c"
	"errors"
)

var bus, err = i2c.Bus(1)

const (
	address = 0x54

	enable_output  = 0x00
	enable_leds    = 0x13
	set_pwm_values = 0x01
	update         = 0x16
)

const (
	Red byte = iota
	Orange
	Yellow
	Green
	Blue
	White
)

func init() {
	if bus != nil || err == nil {
		err = bus.WriteByte(address, enable_output, 0x01)
		err = bus.WriteByteBlock(address, enable_leds, []byte{0xFF, 0xFF, 0xFF})
	}
}

var values = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

var legs = [][]byte{{6, 7, 8, 5, 4, 9}, {17, 16, 15, 13, 11, 10}, {0, 1, 2, 3, 14, 12}}

func PiGlow(led, intensity byte) error {
	if led < 0 || led > 18 {
		return errors.New("Invalid LED")
	}
	if err != nil {
		return err
	}

	values[led] = intensity

	err = bus.WriteByteBlock(address, set_pwm_values, values)
	if err != nil {
		return err
	}
	err = bus.WriteByte(address, update, 0xFF)
	return err
}

func PiGlowLed(leg, color, intensity byte) error {
	if leg < 0 || leg > 2 {
		return errors.New("Invalid leg")
	}
	if color < 0 || color > 5 {
		return errors.New("Invalid color")
	}

	return PiGlow(legs[leg][color], intensity)
}

func PiGlowLeg(leg, intensity byte) error {
	if leg < 0 || leg > 2 {
		return errors.New("Invalid leg")
	}
	if err != nil {
		return err
	}

	for _, led := range legs[leg] {
		values[led] = intensity
	}

	err = bus.WriteByteBlock(address, set_pwm_values, values)
	if err != nil {
		return err
	}
	err = bus.WriteByte(address, update, 0xFF)
	return err
}

func PiGlowRing(color, intensity byte) error {
	if color < 0 || color > 5 {
		return errors.New("Invalid ring")
	}
	if err != nil {
		return err
	}

	values[legs[0][color]] = intensity
	values[legs[1][color]] = intensity
	values[legs[2][color]] = intensity

	err = bus.WriteByteBlock(address, set_pwm_values, values)
	if err != nil {
		return err
	}
	err = bus.WriteByte(address, update, 0xFF)
	return err
}

func ShutDown() error {
	if err != nil {
		return err
	}

	for i := 0; i < 18; i++ {
		values[i] = 0x00
	}

	err = bus.WriteByteBlock(address, set_pwm_values, values)
	if err != nil {
		return err
	}
	err = bus.WriteByte(address, update, 0xFF)

	return err
}
