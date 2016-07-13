package main

import (
	"log"
	"math"
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/ble"
	"github.com/hybridgroup/gobot/platforms/joystick"
)

type pair struct {
	x float64
	y float64
}

type playbackEvent int

const (
	e_takeoff playbackEvent = iota
	e_land
	e_left_x
	e_left_y
	e_right_x
	e_right_y
)

type ev struct {
	t   time.Duration
	p   playbackEvent
	val float64
}

func main() {
	var startTime time.Time
	recording := true
	var events []ev
	event := func(x playbackEvent, val float64) {
		if !recording {
			return
		}
		log.Println("recording event: ", x, val)
		events = append(events, ev{t: time.Now().Sub(startTime), p: x, val: val})
		startTime = time.Now()
	}
	pwd, _ := os.Getwd()
	joystickConfig := pwd + "/dualshock3.json"

	gbot := gobot.NewGobot()

	joystickAdaptor := joystick.NewJoystickAdaptor("ps3")
	joystick := joystick.NewJoystickDriver(joystickAdaptor,
		"ps3",
		joystickConfig,
	)

	droneAdaptor := ble.NewBLEClientAdaptor("ble", os.Args[1])
	drone := ble.NewBLEMinidroneDriver(droneAdaptor, "drone")

	work := func() {

		startTime = time.Now()
		offset := 32767.0
		rightStick := pair{x: 0, y: 0}
		leftStick := pair{x: 0, y: 0}

		gobot.On(joystick.Event("triangle_press"), func(data interface{}) {
			event(e_takeoff, 0)

			drone.TakeOff()
		})
		gobot.On(joystick.Event("x_press"), func(data interface{}) {
			event(e_land, 0)
			drone.Land()
		})
		gobot.On(joystick.Event("circle_press"), func(data interface{}) {
			log.Println("playing back")
			recording = false
			for _, e := range events {
				log.Println("sleepting until next event: ", e.t)
				time.Sleep(e.t)
				switch e.p {
				case e_land:
					log.Println("playing back land")
					drone.Land()
				case e_takeoff:
					log.Println("playing back takeoff")
					drone.TakeOff()
				case e_left_x:
					log.Println("playing back left x = ", e.val)
					leftStick.x = e.val
				case e_left_y:
					log.Println("playing back left y = ", e.val)
					leftStick.y = e.val
				case e_right_x:
					log.Println("playing back right x = ", e.val)
					rightStick.x = e.val
				case e_right_y:
					log.Println("playing back right y = ", e.val)
					rightStick.y = e.val
				}
			}
			drone.Stop()
		})
		gobot.On(joystick.Event("square_press"), func(data interface{}) {
			drone.Stop()
		})

		gobot.On(joystick.Event("left_x"), func(data interface{}) {
			val := float64(data.(int16))
			if leftStick.x != val {
				leftStick.x = val
			}
			event(e_left_x, val)
		})
		gobot.On(joystick.Event("left_y"), func(data interface{}) {
			val := float64(data.(int16))
			if leftStick.y != val {
				leftStick.y = val
			}
			event(e_left_y, val)
		})
		gobot.On(joystick.Event("right_x"), func(data interface{}) {
			val := float64(data.(int16))
			if rightStick.x != val {
				rightStick.x = val
			}
			event(e_right_x, val)
		})
		gobot.On(joystick.Event("right_y"), func(data interface{}) {
			val := float64(data.(int16))
			if rightStick.y != val {
				rightStick.y = val
			}
			event(e_right_y, val)
		})

		gobot.Every(10*time.Millisecond, func() {
			pair := rightStick
			if pair.y < -10 {
				drone.Forward(validatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Backward(validatePitch(pair.y, offset))
			} else {
				drone.Forward(0)
			}

			if pair.x > 10 {
				drone.Right(validatePitch(pair.x, offset))
			} else if pair.x < -10 {
				drone.Left(validatePitch(pair.x, offset))
			} else {
				drone.Right(0)
			}
		})

		gobot.Every(10*time.Millisecond, func() {
			pair := leftStick
			if pair.y < -10 {
				drone.Up(validatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Down(validatePitch(pair.y, offset))
			} else {
				drone.Up(0)
			}

			if pair.x > 20 {
				drone.Clockwise(validatePitch(pair.x, offset))
			} else if pair.x < -20 {
				drone.CounterClockwise(validatePitch(pair.x, offset))
			} else {
				drone.Clockwise(0)
			}
		})
		drone.Stop()
	}

	robot := gobot.NewRobot("minidrone",
		[]gobot.Connection{joystickAdaptor, droneAdaptor},
		[]gobot.Device{joystick, drone},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}

func validatePitch(data float64, offset float64) int {
	value := math.Abs(data) / offset
	if value >= 0.1 {
		if value <= 1.0 {
			return int((float64(int(value*100)) / 100) * 100)
		}
		return 100
	}
	return 0
}
