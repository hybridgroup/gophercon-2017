package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/drivers/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	master := gobot.NewMaster()

	board := edison.NewAdaptor()
	button := gpio.NewGroveButtonDriver(board, "2")
	blue := gpio.NewGroveLedDriver(board, "3")

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
