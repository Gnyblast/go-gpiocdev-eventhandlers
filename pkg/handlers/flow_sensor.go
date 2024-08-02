package handlers

import "github.com/warthog618/go-gpiocdev"

type LiquidFlowSensorHandler[T float32] struct {
	inputMultiplier T
	BaseEventHandler[T]
}

func NewLiquidFlowSensorHandler(inputMultiplier float32) *LiquidFlowSensorHandler[float32] {
	return &LiquidFlowSensorHandler[float32]{
		inputMultiplier: inputMultiplier,
		BaseEventHandler: BaseEventHandler[float32]{
			measurement: 0,
			measured:    make(chan float32),
		},
	}
}

func (h *LiquidFlowSensorHandler[T]) Measure(evt gpiocdev.LineEvent) {
	if evt.Type == gpiocdev.LineEventRisingEdge {
		h.measurement = T(evt.LineSeqno) * h.inputMultiplier
		h.measured <- T(h.measurement)
	}
}

func (h LiquidFlowSensorHandler[T]) GetMeasurement() T {
	return T(h.measurement)
}

func (h LiquidFlowSensorHandler[T]) Subscribe() <-chan T {
	return h.measured
}

func (h *LiquidFlowSensorHandler[T]) CloseChannels() {
	close(h.measured)
}
