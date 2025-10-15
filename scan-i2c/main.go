// I2C Scanner for TinyGo
// Inspired by https://playground.arduino.cc/Main/I2cScanner/
//
// Algorithm
// 1. Send I2C Start condition
// 2. Send a single byte representing the address, and get the ACK/NAK
// 3. Send the stop condition.
// https://electronics.stackexchange.com/a/76620
//
// Learn more about I2C
// https://learn.sparkfun.com/tutorials/i2c/all

package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {

	time.Sleep(time.Second)
	println("Start I2C scanner")
	// Configure I2C using pins specific to the board
	// See https://tinygo.org/docs/reference/microcontrollers/raspberrypi/
	// for the pin mapping of your board.
	//
	// On Raspberry Pi Pico the I2C1 bus is available on multiple pins:
	// - GP2 (SDA) pin4 and GP3 (SCL) pin5
	// - GP6 (SDA) pin9 and GP7 (SCL) pin10
	// - GP10 (SDA) and GP11 (SCL)
	//
	// The default I2C1 pins are GP3 and GP4, so we use those here.
	// Pin 	Hardware pin 	Alternative names
	// GP3 	GPIO3 	I2C1_SCL_PIN
	// GP4 	GPIO4 	I2C1_SDA_PIN
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 100 * machine.KHz,
		SCL: machine.I2C0_SCL_PIN,
		SDA: machine.I2C0_SDA_PIN,
	})

	w := []byte{}
	r := []byte{0} // shall pass at least one byte for I2C code to at all try to communicate
	nDevices := 0

	println("Scanning I2C0...")
	for address := uint16(1); address < 127; address++ {
		if err := machine.I2C0.Tx(address, w, r); err == nil { // try read a byte from the current address
			fmt.Printf("I2C device found at address %#X !\n", address)
			nDevices++
		}
	}

	if nDevices == 0 {
		println("No I2C devices found")
	} else {
		println("Done")
	}

	machine.I2C1.Configure(machine.I2CConfig{
		Frequency: 100 * machine.KHz,
		SCL: machine.I2C1_SCL_PIN,
		SDA: machine.I2C1_SDA_PIN,
	})

	println("Scanning I2C1...")
	for address := uint16(1); address < 127; address++ {
		if err := machine.I2C1.Tx(address, w, r); err == nil { // try read a byte from the current address
			fmt.Printf("I2C device found at address %#X !\n", address)
			nDevices++
		}
	}

	if nDevices == 0 {
		println("No I2C devices found")
	} else {
		println("Done")
	}

}
