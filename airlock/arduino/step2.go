package main

import (
	"fmt"
	"os"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/drivers/gpio"
	"github.com/hybridgroup/gobot/platforms/firmata"
)

func main() {
	master := gobot.NewMaster()
	board := firmata.NewAdaptor(os.Args[1])

	// digital devices
	button := gpio.NewButtonDriver(board, "2")
	blue := gpio.NewLedDriver(board, "3")

	work := func() {
		button.On(button.Event(gpio.ButtonPush), func(data interface{}) {
			fmt.Println("On!")
			blue.On()
		})

		button.On(button.Event(gpio.ButtonRelease), func(data interface{}) {
			fmt.Println("Off!")
			blue.Off()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
