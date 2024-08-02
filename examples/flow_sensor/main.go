package main

import (
	"log"

	"github.com/gnyblast/go-gpiocdev-eventhandlers/pkg/factory"
	"github.com/gnyblast/go-gpiocdev-eventhandlers/pkg/sensors"
	"github.com/warthog618/go-gpiocdev"
	"github.com/warthog618/go-gpiocdev/device/rpi"
)

const inputMultiplier float64 = 0.002222222 // 0.002222222 is the multiplier value for a flow meter that has 450 pulses per liter = (1/450).
const chip string = "gpiochip0"
const pinName string = ""

func main() {

	//Initialize the sensor handler
	lfs := factory.InitializeEventHandlerFor(sensors.FLOW_SENSOR, chip, pinName, inputMultiplier)

	// Get the pin for the device
	pin, err := rpi.Pin(pinName)
	if err != nil {
		log.Fatalf("failed to get pin: %s", err.Error())
	}

	// Get the line by passing the handler from the sensor as WithEventHandler
	line, err := gpiocdev.RequestLine(chip, pin, gpiocdev.WithPullUp, gpiocdev.WithBothEdges, gpiocdev.WithEventHandler(lfs.Measure))
	if err != nil {
		log.Fatalf("failed to request line: %s", err.Error())
	}

	// Close line and channel at the end
	defer line.Close()
	defer lfs.CloseChannels()

	// Listen the measurement if you want to take action or print the values. Do whatever you want.
	// Here shutdown when 100 liters are measured
	for {
		printMeasurement(<-lfs.Subscribe())
		if lfs.GetMeasurement() > 100 {
			break
		}
	}

}

func printMeasurement(measurement float64) {
	log.Printf("%f Liters flowed", measurement)
}
