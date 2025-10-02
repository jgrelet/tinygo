//go:build tinygo && (pico || pico_w || rp2040 || pico2 || pico2_w || rp2350)

package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/ds3231"
)

// Optionnel: injecter une heure à l'édition binaire:
//   tinygo flash ... -ldflags="-X main.buildTime=2025-09-27T06:30:00Z"
var buildTime string // RFC3339, ex: 2006-01-02T15:04:05Z07:00

func must(err error) {
	if err != nil {
		println("ERR:", err.Error())
		panic(err) // stoppe le programme
	}
}

func main() {

	// Initialisation du port série pour le debug
	machine.Serial.Configure(machine.UARTConfig{BaudRate: 115200})
	time.Sleep(2 * time.Second)
	println("DS3231 RTC test started ...")

	// RTC DS3231
	// I2C0 sur GPIO4 (SDA) / GPIO5 (SCL) en 400kHz
	machine.I2C1.Configure(machine.I2CConfig{
		SCL:       machine.I2C1_SCL_PIN,
		SDA:       machine.I2C1_SDA_PIN,
		Frequency: 400 * machine.KHz,
	})

	rtc := ds3231.New(machine.I2C1)
	ok := rtc.Configure()
	if !ok {
		println("DS3231 not detected (addr 0x68) ?")
		for {
			time.Sleep(2 * time.Second)
		}
	}

	// Si l'oscillateur n'est pas en marche, on le démarre.
	if !rtc.IsRunning() {
		must(rtc.SetRunning(true))
	}

	// Mettre à l'heure depuis une valeur injectée (buildTime) :
	if len(buildTime) > 0 && !rtc.IsTimeValid() {
		if t, err := time.Parse(time.RFC3339, buildTime); err == nil {
			must(rtc.SetTime(t.UTC()))
			println("RTC set with buildTime:", t.UTC().Format(time.RFC3339))
		} else {
			println("buildTime invalide, ignore")
		}
	}

	// Affiche l'heure chaque seconde
	for {
		time.Sleep(1 * time.Second)
		// Lire l'heure "RTC"
		t, _ := rtc.ReadTime()
		fmt.Printf("DS3231: %s\n", t.Format("15:04:05 02/01/2006"))
		//fmt.Printf("Temp DS3231: %.2f°C\n", rtc.Temperature())
	}
}