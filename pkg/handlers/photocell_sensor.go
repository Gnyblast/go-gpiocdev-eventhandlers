package handlers

import (
	"time"

	"github.com/warthog618/go-gpiocdev"
	"github.com/warthog618/go-gpiocdev/device/rpi"
)

type PhotocellSensorHandler[T int] struct {
	ongoingMeasurement int
	lineChip           string
	pinNumber          string
	BaseEventHandler[T]
}

func NewPhotocellSensorHandler(chip string, pin string) *PhotocellSensorHandler[int] {

	return &PhotocellSensorHandler[int]{
		ongoingMeasurement: 0,
		lineChip:           chip,
		pinNumber:          pin,
		BaseEventHandler: BaseEventHandler[int]{
			measurement: 0,
			measured:    make(chan int),
		},
	}
}

func (h *PhotocellSensorHandler[T]) Measure(evt gpiocdev.LineEvent) {
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
		h.measurement = T(h.ongoingMeasurement)
		h.ongoingMeasurement = 0
		h.measured <- h.measurement
	}
}

func (h PhotocellSensorHandler[T]) setLineToZero() error {
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

func (h PhotocellSensorHandler[T]) setLineToInput() error {
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

func (h PhotocellSensorHandler[T]) GetMeasurement() T {
	return T(h.measurement)
}

func (h PhotocellSensorHandler[T]) Subscribe() <-chan T {
	return h.measured
}

func (h *PhotocellSensorHandler[T]) CloseChannels() {
	close(h.measured)
}
