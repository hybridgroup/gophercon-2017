package main

import (
	"fmt"
	"os"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/drivers/gpio"
	"github.com/hybridgroup/gobot/platforms/firmata"
)

var button *gpio.ButtonDriver
var blue *gpio.LedDriver
var green *gpio.LedDriver

func TurnOff() {
	blue.Off()
	green.Off()
}

func Reset() {
	TurnOff()
	fmt.Println("Airlock ready.")
	green.On()
}

func main() {
	master := gobot.NewMaster()

	board := firmata.NewAdaptor(os.Args[1])

	// digital devices
	button = gpio.NewButtonDriver(board, "2")
	blue = gpio.NewLedDriver(board, "3")
	green = gpio.NewLedDriver(board, "4")

	work := func() {
		Reset()

		button.On(button.Event(gpio.ButtonPush), func(data interface{}) {
			TurnOff()
			fmt.Println("On!")
			blue.On()
		})

		button.On(button.Event(gpio.ButtonRelease), func(data interface{}) {
			Reset()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
