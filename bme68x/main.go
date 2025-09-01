package main

import (
	"fmt"
	"log"
	"machine"
	"time"
    "BME68x/bme68x"
    "strings"
)

func main() {

	const seaLevelPressurehPa = 1013.25
	const humidityDescription = "Comfortable"
	machine.I2C1.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
		SCL: machine.I2C1_SCL_PIN,
		SDA: machine.I2C1_SDA_PIN,
	})

	time.Sleep(time.Second)
	println("Start sampling BME860 sensor")

	tsensor := bme68x.NewI2C(machine.I2C1,
		bme68x.WithIIRFilter(bme68x.Coeff4),
		bme68x.WithTemperatureOversampling(bme68x.Sampling8X),
		bme68x.WithPressureOversampling(bme68x.Sampling4X),
		bme68x.WithHumidityOversampling(bme68x.Sampling2X),
		bme68x.WithHeatrDuration(150),
		bme68x.WithHeatrTemperature(320),
	)
	if err := tsensor.Configure(); err != nil {
		log.Fatal(fmt.Sprintf("Fatal configuring sensor: %s", err))
		return
	}

	connected, err := tsensor.Connected()
	if err != nil {
		log.Fatal(fmt.Sprintf("Fatal checking sensor connection: %s", err))
		return
	}

	if !connected {
		log.Fatal("sensor not connected")
		return
	}

	if err := tsensor.SetMode(bme68x.ModeForced); err != nil {
		log.Fatal(fmt.Sprintf("Fatal setting sensor mode: %s", err))
		return
	}

	for {
		if err := tsensor.Read(); err != nil {
			log.Fatal(fmt.Sprintf("Fatal reading sensor: %s", err))

			time.Sleep(2 * time.Second)
			continue
		}

		log.Print(strings.Repeat("-", 40))

		log.Print(fmt.Sprintf("    Temperature: %.2f°C", tsensor.Temperature))
		log.Print(fmt.Sprintf("    Pressure: %.fhPa", tsensor.Pressure/100))
		log.Print(fmt.Sprintf("    Gas: %.1fKOhms", tsensor.GasResistance/1000))
		log.Print(fmt.Sprintf("    Approx. Altitude: %.1fm", bme68x.CalcAltitude(seaLevelPressurehPa, tsensor.Pressure)))
		log.Print(fmt.Sprintf("    Humidity: %.1f%% (%s)", tsensor.Humidity, humidityDescription))
		log.Print(strings.Repeat("-", 40))

		time.Sleep(2 * time.Second)
	}
}
