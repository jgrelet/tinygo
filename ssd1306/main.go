//from https://github.com/Nondzu/ssd1306_font

package main

import (
	"fmt"
	"machine"
	"time"

	font "github.com/Nondzu/ssd1306_font"
	"tinygo.org/x/drivers/ssd1306"
)

func main() {

	var (
		with int16 = 128
		//height int16 = 32
		height int16 = 64
	)
	time.Sleep(time.Second)
	println("Start Oled display") // Please wait some time after turning on the device to properly initialize the display

	// The default I2C1 pins are GP3 and GP4, so we use those here.
	machine.I2C1.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
		SCL:       machine.I2C1_SCL_PIN,
		SDA:       machine.I2C1_SDA_PIN,
	})

	// Init OLED Display
	dev := ssd1306.NewI2C(machine.I2C1)
	dev.Configure(ssd1306.Config{
		Width:    with,
		Height:   height,
		Address:  0x3C,
		VccState: ssd1306.SWITCHCAPVCC,
	})
	dev.ClearBuffer()
	dev.ClearDisplay()

	// Init font library
	display := font.NewDisplay(dev)
	display.Configure(font.Config{FontType: font.FONT_7x10}) //set font here

	i := 0
	for {
		display.YPos = 0                               // set position Y
		display.PrintText(fmt.Sprintf("Count: %d", i)) // print text
		//display.PrintText("Hello World!" + i) // print text
		display.YPos = 12                   // set position Y
		display.XPos = 0                    // set position X
		display.PrintText("Temp = 20 C")    // print text
		display.YPos = 24                   // set position Y
		display.XPos = 0                    // set position X
		display.PrintText("Pres = 1023 mb") // print text
		time.Sleep(time.Second * 1)
		i = i + 1

	}
}
