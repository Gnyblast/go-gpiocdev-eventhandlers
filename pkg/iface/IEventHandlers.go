package iface

import "github.com/warthog618/go-gpiocdev"

type IEventHandlers interface {
	Measure(evt gpiocdev.LineEvent)
	GetMeasurement() float64
	Subscribe() <-chan float64
	CloseChannels()
}
