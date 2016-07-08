package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/firmata"
)

func main() {
	gbot := gobot.NewGobot()

	board := firmata.NewFirmataAdaptor("arduino", os.Args[1])

	// digital devices
	button := gpio.NewGroveButtonDriver(board, "button", "2")
	blue := gpio.NewGroveLedDriver(board, "blue", "3")

	work := func() {
		gobot.On(button.Event(gpio.Push), func(data interface{}) {
			fmt.Println("On!")
			blue.On()
		})

		gobot.On(button.Event(gpio.Release), func(data interface{}) {
			fmt.Println("Off!")
			blue.Off()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
