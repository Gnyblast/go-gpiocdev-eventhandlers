package handlers

type BaseEventHandler struct {
	measurement float64
	measured    chan float64
}

func (b BaseEventHandler) GetMeasurement() float64 {
	return b.measurement
}

func (b BaseEventHandler) Subscribe() <-chan float64 {
	return b.measured
}

func (h *BaseEventHandler) CloseChannels() {
	close(h.measured)
}
