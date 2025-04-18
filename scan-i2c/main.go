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
	machine.I2C1.Configure(machine.I2CConfig{
		Frequency: 100 * machine.KHz,
		//SCL: machine.I2C1_SCL_PIN,
		//SDA: machine.I2C1_SDA_PIN,
	})

	w := []byte{}
	r := []byte{0} // shall pass at least one byte for I2C code to at all try to communicate
	nDevices := 0

	println("Scanning...")
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

	// procrastinate for an hour to ensure everything was printed out and board does not die
	time.Sleep(1 * time.Hour)

}
