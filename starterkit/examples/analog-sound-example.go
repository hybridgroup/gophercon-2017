package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("COM111")
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, "0")

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			sound := data.(int)
			sound = sound / 10.0
			fmt.Println("sound", sound)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor},
		work,
	)

	robot.Start()
}
