// blinky.go

// tinygo flash -target=pico-w -monitor ./main.go
// or
// make run

package main

import (
	"machine"
	"time"
	"fmt"
)

func main() {
	time.Sleep(time.Second)
	println("Start blinking")
	//led := machine.LED
	led := machine.GP15 // pin 20
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	for {
		led.High()
		fmt.Printf("Led on: %t\n", led.Get())
		time.Sleep(1 * time.Second)
		led.Low()
		fmt.Printf("Led off: %t\n", led.Get())
		time.Sleep(1 * time.Second)
	}
}
