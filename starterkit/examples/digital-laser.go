package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("COM111")
	laser := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		laser.On()

	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{laser},

		work,
	)

	robot.Start()

}
