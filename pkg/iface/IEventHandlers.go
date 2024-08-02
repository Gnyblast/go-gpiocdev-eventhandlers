package iface

import "github.com/warthog618/go-gpiocdev"

type IEventHandlers[T any] interface {
	Measure(evt gpiocdev.LineEvent)
	GetMeasurement() T
	Subscribe() <-chan T
	CloseChannels()
}
