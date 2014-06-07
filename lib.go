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

var values = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

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
