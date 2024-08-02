package main

import (
	"log"

	"github.com/gnyblast/go-gpiocdev-eventhandlers/pkg/handlers"
	"github.com/warthog618/go-gpiocdev"
	"github.com/warthog618/go-gpiocdev/device/rpi"
)

const chip string = "gpiochip0"
const pinName string = "GPIO4"

func main() {

	//Initialize the sensor handler, this handler wants chipname and pinname to be able to set it to output as 0 then start listening.
	pcs := handlers.NewPhotocellSensorHandler(chip, pinName)

	// Get the pin for the device
	pin, err := rpi.Pin(pinName)
	if err != nil {
		log.Fatalf("failed to get pin: %s", err.Error())
	}

	// Get the line by passing the handler from the sensor as WithEventHandler
	line, err := gpiocdev.RequestLine(chip, pin, gpiocdev.WithPullUp, gpiocdev.WithBothEdges, gpiocdev.WithEventHandler(pcs.Measure))
	if err != nil {
		log.Fatalf("failed to request line: %s", err.Error())
	}

	// Close line and channel at the end
	defer line.Close()
	defer pcs.CloseChannels()

	// listen the measurement if you want to take action or print the values. Do whatever you want.
	for {
		printMeasurement(<-pcs.Subscribe())
		if pcs.GetMeasurement() > 100 {
			break
		}
	}

}

func printMeasurement(measurement int) {
	log.Printf("Brightness is: %d", measurement)
}
