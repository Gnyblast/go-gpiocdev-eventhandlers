package factory

import (
	"github.com/gnyblast/go-gpiocdev-eventhandlers/pkg/handlers"
	"github.com/gnyblast/go-gpiocdev-eventhandlers/pkg/iface"
	"github.com/gnyblast/go-gpiocdev-eventhandlers/pkg/sensors"
)

func InitializeEventHandlerFor(sensorType sensors.Sensors, sensorChip string, sensorPin string, inputMultipler float64) iface.IEventHandlers {

	if sensorType == sensors.FLOW_SENSOR {
		return handlers.NewLiquidFlowSensorHandler(inputMultipler)
	}

	if sensorType == sensors.PHOTOCELL_SENSOR {
		return handlers.NewPhotocellSensorHandler(sensorChip, sensorPin)
	}

	return handlers.NewLiquidFlowSensorHandler(inputMultipler)
}
