package iface

import "github.com/warthog618/go-gpiocdev"

type IHandler[T any] interface {
	Measure(evt gpiocdev.LineEvent)
	GetMeasurement() T
	Subscribe() <-chan T
	CloseChannels()
}
