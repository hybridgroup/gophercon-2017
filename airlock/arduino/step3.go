package main

import (
	"fmt"
	"os"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
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
	gbot := gobot.NewGobot()

	board := firmata.NewFirmataAdaptor("arduino", os.Args[1])

	// digital devices
	button = gpio.NewButtonDriver(board, "button", "2")
	blue = gpio.NewLedDriver(board, "blue", "3")
	green = gpio.NewLedDriver(board, "green", "4")

	work := func() {
		Reset()

		gobot.On(button.Event(gpio.Push), func(data interface{}) {
			TurnOff()
			fmt.Println("On!")
			blue.On()
		})

		gobot.On(button.Event(gpio.Release), func(data interface{}) {
			Reset()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
