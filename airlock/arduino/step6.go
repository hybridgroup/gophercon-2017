package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/drivers/gpio"
	"github.com/hybridgroup/gobot/platforms/firmata"
)

var button *gpio.ButtonDriver
var blue *gpio.LedDriver
var green *gpio.LedDriver
var buzzer *gpio.GroveBuzzerDriver
var touch *gpio.ButtonDriver
var light *gpio.AnalogSensorDriver

func Doorbell() {
	TurnOff()
	blue.On()
	buzzer.Tone(gpio.C4, gpio.Quarter)
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
	master := gobot.NewMaster()

	a := api.NewAPI(master)
	a.Start()

	board := firmata.NewAdaptor(os.Args[1])

	// digital devices
	button = gpio.NewButtonDriver(board, "2")
	blue = gpio.NewLedDriver(board, "3")
	green = gpio.NewLedDriver(board, "4")
	buzzer = gpio.NewGroveBuzzerDriver(board, "7")
	touch = gpio.NewButtonDriver(board, "8")

	// analog devices
	light = gpio.NewAnalogSensorDriver(board, "0")

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

		touch.On(touch.Event(gpio.ButtonPush), func(data interface{}) {
			Doorbell()
		})

		light.On(light.Event("data"), func(data interface{}) {
			fmt.Println("light", data)
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green, buzzer, touch, light},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
