//go:build tinygo && (pico || pico_w || rp2040 || pico2 || pico2_w || rp2350)

package main

import (
	"fmt"
	"machine"
	"time"

	font "github.com/Nondzu/ssd1306_font"
	"tinygo.org/x/drivers/ssd1306"
	//"github.com/jgrelet/pico-rtc/ssd1306x"
	"tinygo.org/x/drivers/ds3231"
	ntp "github.com/jgrelet/pico-rtc/ntputil"
)


// mustRetry attempts to execute the provided function f up to n times, waiting for the specified delay between attempts.
// If f returns nil, mustRetry returns immediately. If all attempts fail, it prints the error and panics with the last error.
func mustRetry(n int, delay time.Duration, f func() error) {
	for i := 0; i < n; i++ {
		if err := f(); err == nil {
			return
		} else if i == n-1 {
			println("ERR:", err.Error())
			panic(err)
		}
		time.Sleep(delay)
	}
}


// main initializes and runs the DS3231 NTP + RTC test application.
// 
// This function performs the following tasks:
//   - Configures the serial port for debugging output.
//   - Initializes the I2C interface and configures the SSD1306 OLED display.
//   - Sets up the font library for text rendering on the display.
//   - Configures the I2C interface for the DS3231 RTC module.
//   - Establishes a Wi-Fi connection and synchronizes time with an NTP server.
//   - Initializes the DS3231 RTC, ensuring the oscillator is running and sets its time to the NTP time.
//   - Enters a loop where it reads the current time and temperature from the DS3231 every second,
//     displaying this information on the OLED display and printing it to the serial output.
//
// Any errors encountered during initialization or operation are reported via the serial output and the OLED display.
func main() {

	var (
		with int16 = 128
		//height int16 = 32
		height int16 = 64
	)

	// Initialisation du port série pour le debug
	machine.Serial.Configure(machine.UARTConfig{BaudRate: 115200})
	time.Sleep(2 * time.Second)
	println("DS3231 NTP + RTC test started ...")

	// --- OLED ---
	// The default I2C1 pins are GP3 and GP4, so we use those here.
	machine.I2C1.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
		SCL: machine.I2C1_SCL_PIN,
		SDA: machine.I2C1_SDA_PIN,
	})
	//machine.I2C1.Configure(machine.I2CConfig{Frequency: 400000})

	// Display
	dev := ssd1306.NewI2C(machine.I2C1)
	dev.Configure(ssd1306.Config{Width: with, Height: height, Address: 0x3C, VccState: ssd1306.SWITCHCAPVCC})
	dev.ClearBuffer()
	dev.ClearDisplay()

	//font library init
	display := font.NewDisplay(*dev)
	display.Configure(font.Config{FontType: font.FONT_7x10}) //set font here
	//disp := &ssd1306.Display{dev: *dev, width: with, height: height}
	display.YPos = 0                                         // set position Y
	display.XPos = 0   

	/* 
	// --- OLED ---
	disp := ssd1306x.NewI2C(ssd1306x.Config{
		I2C:     *machine.I2C1,
		Address: 0x3C,
		SCL:     machine.I2C1_SCL_PIN, // Pico/Pico2: GP5
		SDA:     machine.I2C1_SDA_PIN, // Pico/Pico2: GP4
		Freq:    400 * machine.KHz,
		Width:   128,
		Height:  64,
	})

	//font library init
	display := font.NewDisplay(*disp.Device())               //pass by value
	display.Configure(font.Config{FontType: font.FONT_7x10}) //set font here 
	*/

	// RTC DS3231
	// I2C0 sur GPIO4 (SDA) / GPIO5 (SCL) en 400kHz
	machine.I2C0.Configure(machine.I2CConfig{
		SCL:       machine.I2C0_SCL_PIN,
		SDA:       machine.I2C0_SDA_PIN,
		Frequency: 400 * machine.KHz,
	})

	// Initialiser le Wi-Fi et la connexion NTP
	conn, err := ntp.NewNTPConn("Pico2-w", "192.168.1.149", 10, /*logger.Logger*/ nil)
	if err != nil {
		fmt.Println("Error connect Wi-Fi :", err)
		display.PrintText(fmt.Sprintf("Error Wi-Fi:", err))
		dev.Display()
		return
	}
	//logger.Logger.Info(conn.String())
	println(conn.String())

	now, err := conn.GetNTPTime()
	if err != nil {
		fmt.Println("NTP error:", err)
		display.PrintText(fmt.Sprintf("NTP error:", err))
		dev.Display()
	} else {
		//logger.Logger.Info("NTP time :", now.String())
		fmt.Println("NTP time : ", now.String())
	}

	// Initialiser le module RTC DS3231
	// Adresse I2C0 0x68, pin 6 GP4 (SDA) / pin 7 GP5 (SCL) en 400kHz
	machine.I2C0.Configure(machine.I2CConfig{
		SCL:       machine.I2C0_SCL_PIN,
		SDA:       machine.I2C0_SDA_PIN,
		Frequency: 400 * machine.KHz,
	})
	rtc := ds3231.New(machine.I2C0)
	ok := rtc.Configure()
	if !ok {
		println("DS3231 not detected (addr 0x68) ?")
		for {
			time.Sleep(1 * time.Second)
		}
	}

	// Si l'oscillateur n'est pas en marche, on le démarre.
	if !rtc.IsRunning() {
		//check(rtc.SetRunning(true))
		println("DS3231 was not running, starting it ...")
		mustRetry(5, 200*time.Millisecond, func() error { return rtc.SetRunning(true) })

	}

	mustRetry(5, 200*time.Millisecond, func() error { return rtc.SetTime(now) })
	println("DS3231 time set to NTP time")

	// Affiche l'heure chaque seconde
	for {
		time.Sleep(1 * time.Second)
		// Lire l'heure "RTC"
		t, err := rtc.ReadTime()
		if err != nil {
			println("DS3231 ReadTime error:", err.Error())
			continue
		}
		temp, err := rtc.ReadTemperature()
		if err != nil {
			println("DS3231 ReadTemperature error:", err.Error())
			continue
		}
		T := float32(temp)/1000.0 // en °C
		// Afficher l'heure
		//fmt.Printf("DS3231: %s\n", t.Format("15:04:05 02/01/2006"))
		// Afficher l'heure et la température
		fmt.Printf("DS3231: %s, Temp: %3.0f°C\n", t.Format("15:04:05 02/01/2006"), T)
		display.YPos = 0
		display.PrintText(t.Format("15:04:05 02/01/06"))
		display.YPos = 12
		display.PrintText(fmt.Sprintf("Temp: %2.0f C", T))
		dev.Display()
		dev.ClearBuffer()
	}
}