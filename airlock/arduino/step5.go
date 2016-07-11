package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

var button *gpio.ButtonDriver
var blue *gpio.LedDriver
var green *gpio.LedDriver
var buzzer *gpio.GroveBuzzerDriver
var touch *gpio.ButtonDriver

func Doorbell() {
	TurnOff()
	blue.On()
	buzzer.Tone(gpio.C4, gpio.Half)
	<-time.After(1 * time.Second)
	Reset()
}

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

	a := api.NewAPI(gbot)
	a.Start()

	board := firmata.NewFirmataAdaptor("arduino", os.Args[1])

	// digital devices
	button = gpio.NewButtonDriver(board, "button", "2")
	blue = gpio.NewLedDriver(board, "blue", "3")
	green = gpio.NewLedDriver(board, "green", "4")
	buzzer = gpio.NewGroveBuzzerDriver(board, "buzzer", "7")
	touch = gpio.NewButtonDriver(board, "touch", "8")

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

		gobot.On(touch.Event(gpio.Push), func(data interface{}) {
			Doorbell()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green, buzzer, touch},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
