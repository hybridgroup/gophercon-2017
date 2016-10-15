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
var red *gpio.GroveLedDriver
var buzzer *gpio.GroveBuzzerDriver
var touch *gpio.GroveTouchDriver
var rotary *gpio.GroveRotaryDriver
var sensor *gpio.GroveTemperatureSensorDriver

func CheckFireAlarm() {
	temp := sensor.Temperature()
	fmt.Println("Current temperature:", temp)
	if temp >= 15 {
		TurnOff()
		red.On()
		buzzer.Tone(gpio.F4, gpio.Half)
	}
}

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

	board := edison.NewAdaptor()

	// digital
	button = gpio.NewGroveButtonDriver(board, "2")
	blue = gpio.NewGroveLedDriver(board, "3")
	green = gpio.NewGroveLedDriver(board, "4")
	red = gpio.NewGroveLedDriver(board, "5")
	buzzer = gpio.NewGroveBuzzerDriver(board, "6")
	touch = gpio.NewGroveTouchDriver(board, "8")

	// analog
	rotary = gpio.NewGroveRotaryDriver(board, "0")
	sensor = gpio.NewGroveTemperatureSensorDriver(board, "1")

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

		gobot.Every(1*time.Second, func() {
			CheckFireAlarm()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green, red, buzzer, touch, rotary, sensor},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
