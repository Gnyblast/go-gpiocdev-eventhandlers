package handlers

import (
	"time"

	"github.com/warthog618/go-gpiocdev"
	"github.com/warthog618/go-gpiocdev/device/rpi"
)

type PhotocellSensorHandler struct {
	ongoingMeasurement float64
	lineChip           string
	pinNumber          string
	BaseEventHandler
}

func NewPhotocellSensorHandler(chip string, pin string) *PhotocellSensorHandler {

	return &PhotocellSensorHandler{
		ongoingMeasurement: 0,
		lineChip:           chip,
		pinNumber:          pin,
		BaseEventHandler: BaseEventHandler{
			measurement: 0,
			measured:    make(chan float64),
		},
	}
}

func (h *PhotocellSensorHandler) Measure(evt gpiocdev.LineEvent) {
	if h.ongoingMeasurement == 0 {
		err := h.setLineToZero()
		if err != nil {
			h.measurement = 0
			h.measured <- 0
			return
		}

		time.Sleep(100 * time.Millisecond)
		err = h.setLineToInput()
		if err != nil {
			h.measurement = 0
			h.measured <- 0
			return
		}
	}

	if evt.Type == gpiocdev.LineEventFallingEdge {
		h.ongoingMeasurement++
		return
	}

	if evt.Type == gpiocdev.LineEventRisingEdge {
		h.measurement = h.ongoingMeasurement
		h.ongoingMeasurement = 0
		h.measured <- h.measurement
	}
}

func (h PhotocellSensorHandler) setLineToZero() error {
	pin, err := rpi.Pin(h.pinNumber)
	if err != nil {
		return err
	}

	line, err := gpiocdev.RequestLine(h.lineChip, pin, gpiocdev.AsOutput(0))
	if err != nil {
		return err
	}
	defer line.Close()

	return nil
}

func (h PhotocellSensorHandler) setLineToInput() error {
	pin, err := rpi.Pin(h.pinNumber)
	if err != nil {
		return err
	}

	line, err := gpiocdev.RequestLine(h.lineChip, pin)
	if err != nil {
		return err
	}
	defer line.Close()

	err = line.Reconfigure(gpiocdev.AsInput)
	if err != nil {
		return err
	}

	return nil
}
