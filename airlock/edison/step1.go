package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/drivers/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	master := gobot.NewMaster()

	board := edison.NewAdaptor()
	blue := gpio.NewGroveLedDriver(board, "3")

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

	master.AddRobot(robot)

	master.Start()
}
