# gpiocdev-eventhandler

This repository aims to provide handlers for the [go-gpicdev](https://github.com/warthog618/go-gpiocdev) Library's `gpiocdev.WithEventHandler` line option.

It's still under development with two sensor added to the library without tests are performed.

## Usage

- Each handler is a stuct that implements `IEventHandlers` generic interface. 
- `Measure()` method is used as the main handler that does the logic/calculation/measurement for the sensor data. This should be passed to `WithEeventHandler()` method.
- `Subscribe()` is for listening measurement if immediate responses are needed when a measurement is done. It's a channel type.
- `CloseChannels()` is the method that closes the above channel. Best to be called when usage is no more required.
- `GetMeasurement()` is the method that return the latest measurement.

A simple example might look like the following:

```go
lfs := handlers.NewLiquidFlowSensorHandler(0.002222222)
defer lfs.CloseChannels()

gpiocdev.RequestLine("gpiochip0", pin, gpiocdev.WithPullUp, gpiocdev.WithBothEdges, gpiocdev.WithEventHandler(lfs.Measure))

for {
    fmt.Printf("%f Liters Flowed",<-lfs.Subscribe())
}
```

Please refer to the [Examples](#examples) section to see more details examples for each sensor.

## Examples

|Sensor|Image|Example|
|------|-----|-------|
|Water Flow Sensor|![WaterFlowSensor](/images/water_flow_sensor.jpg)|[example](/examples/flow_sensor/main.go)|
|Photocell Sensor|![Photocell](/images/photocell.jpg)|[example](/examples/photocell_sensor/main.go)|
