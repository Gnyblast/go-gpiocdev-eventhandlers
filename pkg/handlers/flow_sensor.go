package handlers

import "github.com/warthog618/go-gpiocdev"

type LiquidFlowSensorHandler struct {
	inputMultiplier float64
	BaseEventHandler
}

func NewLiquidFlowSensorHandler(inputMultiplier float64) *LiquidFlowSensorHandler {
	return &LiquidFlowSensorHandler{
		inputMultiplier: inputMultiplier,
		BaseEventHandler: BaseEventHandler{
			measurement: 0,
			measured:    make(chan float64),
		},
	}
}

func (h *LiquidFlowSensorHandler) Measure(evt gpiocdev.LineEvent) {
	if evt.Type == gpiocdev.LineEventRisingEdge {
		h.measurement = float64(evt.LineSeqno) * h.inputMultiplier
		h.measured <- h.measurement
	}
}
