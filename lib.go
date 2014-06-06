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

func init() {
	if bus != nil || err == nil {
		err = bus.WriteByte(address, enable_output, 0x01)
		err = bus.WriteByteBlock(address, enable_leds, []byte{0xFF, 0xFF, 0xFF})
	}
}

var values = []byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x18, 0x20, 0x30, 0x40, 0x50, 0x60, 0x70, 0x80, 0x90, 0xA0, 0xC0, 0xE0, 0xFF}

func PiGlow(led, intensity byte) error {
	if led < 0 || led > 18 {
		return errors.New("Invalid LED")
	}
	if err != nil {
		return err
	}

	values[led] = intensity

	bus.WriteByteBlock(address, set_pwm_values, values)
	bus.WriteByte(address, update, 0xFF)

	//bus.WriteByteBlock(address, reg, list)

	return nil
}
