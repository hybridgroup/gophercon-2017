package main

import (
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	board := firmata.NewAdaptor(os.Args[1])
	blue := gpio.NewLedDriver(board, "3")

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

	robot.Start()
}
