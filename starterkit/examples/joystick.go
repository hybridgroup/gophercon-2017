package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("COM43")
	xAxis := aio.NewAnalogSensorDriver(firmataAdaptor, "0")
	yAxis := aio.NewAnalogSensorDriver(firmataAdaptor, "1")
	button := gpio.NewButtonDriver(firmataAdaptor, "13")

	work := func() {
		xAxis.On(aio.Data, func(data interface{}) {
			xAxis := data.(int)
			xAxis = xAxis / 10.0
			fmt.Println("x axis", xAxis)

		})

		yAxis.On(aio.Data, func(data interface{}) {
			yAxis := data.(int)
			yAxis = yAxis / 10.0
			fmt.Println("y axis", yAxis)

		})

		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pressed")
		})
		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button released")
		})

	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{xAxis, yAxis, button},
		work,
	)

	robot.Start()
}
