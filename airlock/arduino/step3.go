package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
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
	board := firmata.NewAdaptor(os.Args[1])

	// digital devices
	button = gpio.NewButtonDriver(board, "2")
	blue = gpio.NewLedDriver(board, "3")
	green = gpio.NewLedDriver(board, "4")

	work := func() {
		Reset()

		button.On(gpio.ButtonPush, func(data interface{}) {
			TurnOff()
			fmt.Println("On!")
			blue.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			Reset()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green},
		work,
	)

	robot.Start()
}
