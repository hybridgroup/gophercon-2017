package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/gpio"
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
var sound *gpio.GroveSoundSensorDriver
var light *gpio.GroveLightSensorDriver

func DetectSound(level int) {
	if level >= 400 {
		fmt.Println("Sound detected")
		TurnOff()
		blue.On()
		<-time.After(1 * time.Second)
		Reset()
	}
}

func DetectLight(level int) {
	if level >= 400 {
		fmt.Println("Light detected")
		TurnOff()
		blue.On()
		<-time.After(1 * time.Second)
		Reset()
	}
}

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
	gbot := gobot.NewGobot()

	a := api.NewAPI(gbot)
	a.Start()

	// digital
	board := edison.NewEdisonAdaptor("edison")

	button = gpio.NewGroveButtonDriver(board, "button", "2")
	blue = gpio.NewGroveLedDriver(board, "blue", "3")
	green = gpio.NewGroveLedDriver(board, "green", "4")
	red = gpio.NewGroveLedDriver(board, "red", "5")
	buzzer = gpio.NewGroveBuzzerDriver(board, "buzzer", "7")
	touch = gpio.NewGroveTouchDriver(board, "touch", "8")

	// analog
	rotary = gpio.NewGroveRotaryDriver(board, "rotary", "0")
	sensor = gpio.NewGroveTemperatureSensorDriver(board, "sensor", "1")
	sound = gpio.NewGroveSoundSensorDriver(board, "sound", "2")
	light = gpio.NewGroveLightSensorDriver(board, "light", "3")

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

		gobot.On(rotary.Event("data"), func(data interface{}) {
			fmt.Println("rotary", data)
		})

		gobot.On(sound.Event("data"), func(data interface{}) {
			DetectSound(data.(int))
		})

		gobot.On(light.Event("data"), func(data interface{}) {
			DetectLight(data.(int))
		})

		gobot.Every(1*time.Second, func() {
			CheckFireAlarm()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green, red, buzzer, touch, rotary, sensor, sound, light},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
