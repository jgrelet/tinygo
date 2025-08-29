//from https://github.com/Nondzu/ssd1306_font

package main

import (
	"machine"
	"time"

	font "github.com/Nondzu/ssd1306_font"
	"tinygo.org/x/drivers/ssd1306"
)

func main() {

	time.Sleep(time.Second)
	println("Start Oled display") // Please wait some time after turning on the device to properly initialize the display
	machine.I2C1.Configure(machine.I2CConfig{
		//Frequency: 400000,
		Frequency: 400 * machine.KHz,
		SCL: machine.I2C1_SCL_PIN,
		SDA: machine.I2C1_SDA_PIN,
	})
	
	//machine.I2C1.Configure(machine.I2CConfig{Frequency: 400000})

	// Display
	dev := ssd1306.NewI2C(machine.I2C1)
	dev.Configure(ssd1306.Config{Width: 128, Height: 64, Address: 0x3C, VccState: ssd1306.SWITCHCAPVCC})
	dev.ClearBuffer()
	dev.ClearDisplay()

	//font library init
	display := font.NewDisplay(dev)
	display.Configure(font.Config{FontType: font.FONT_7x10}) //set font here

	display.YPos = 0                 // set position Y
	display.XPos = 0                  // set position X
	display.PrintText("Hello World!") // print text
	display.YPos = 12                // set position Y
	display.XPos = 0                  // set position X
	display.PrintText("Temp = 20 C") // print text
	display.YPos = 24               // set position Y
	display.XPos = 0                  // set position X
	display.PrintText("Pres = 1023 mb") // print text      

	for {
	}
}