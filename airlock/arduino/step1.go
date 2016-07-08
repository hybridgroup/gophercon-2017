package main

import (
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	board := firmata.NewFirmataAdaptor("arduino", os.Args[1])

	// digital devices
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
