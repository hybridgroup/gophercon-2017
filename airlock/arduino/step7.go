package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var button *gpio.GroveButtonDriver
var blue *gpio.GroveLedDriver
var green *gpio.GroveLedDriver
var red *gpio.GroveLedDriver
var buzzer *gpio.GroveBuzzerDriver
var touch *gpio.GroveTouchDriver
var light *aio.AnalogSensorDriver
var rotary *aio.GroveRotaryDriver
var sensor *aio.GroveTemperatureSensorDriver

func CheckFireAlarm() {
	temp := sensor.Temperature()
	fmt.Println("Current temperature:", temp)
	if temp >= 40 {
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

	board := firmata.NewAdaptor(os.Args[1])

	// digital devices
	button = gpio.NewGroveButtonDriver(board, "2")
	blue = gpio.NewGroveLedDriver(board, "3")
	green = gpio.NewGroveLedDriver(board, "4")
	red = gpio.NewGroveLedDriver(board, "5")
	buzzer = gpio.NewGroveBuzzerDriver(board, "7")
	touch = gpio.NewGroveTouchDriver(board, "8")

	// analog devices
	rotary = aio.NewGroveRotaryDriver(board, "0")
	sensor = aio.NewGroveTemperatureSensorDriver(board, "1")

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

		button.On(gpio.ButtonPush, func(data interface{}) {
			Doorbell()
		})

		rotary.On(aio.Data, func(data interface{}) {
			fmt.Println("rotary", data)
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
