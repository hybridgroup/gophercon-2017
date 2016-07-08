package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	board := edison.NewEdisonAdaptor("edison")
	blue := gpio.NewGroveLedDriver(board, "blue", "3")

	work := func() {
		gobot.Every(1*time.Second, func() {
			blue.Toggle()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{blue},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
