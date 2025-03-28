package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// button := machine.D2 // Arduino Uno
	button := machine.GP22  // Raspberry Pi Pico
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	println("Press the button!")

	for {
		// if the button is low (pressed)
		if !button.Get() {
			// toggle the LED
			led.Set(!led.Get())
		}

		// wait a bit, for the blinking effect
		time.Sleep(200 * time.Millisecond)
	}
}