package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/drivers/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

var button *gpio.GroveButtonDriver
var blue *gpio.GroveLedDriver
var green *gpio.GroveLedDriver
var buzzer *gpio.GroveBuzzerDriver
var touch *gpio.GroveTouchDriver

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
	master := gobot.NewMaster()

	a := api.NewAPI(master)
	a.Start()

	board := edison.NewAdaptor()
	button = gpio.NewGroveButtonDriver(board, "2")
	blue = gpio.NewGroveLedDriver(board, "3")
	green = gpio.NewGroveLedDriver(board, "4")
	buzzer = gpio.NewGroveBuzzerDriver(board, "6")
	touch = gpio.NewGroveTouchDriver(board, "8")

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
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green, buzzer, touch},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
