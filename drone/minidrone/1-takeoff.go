package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/ble"
)

func main() {
	master := gobot.NewMaster()

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	drone := ble.NewMinidroneDriver(bleAdaptor)

	work := func() {
		fmt.Println("takeoff...")
		drone.TakeOff()

		gobot.After(10*time.Second, func() {
			fmt.Println("landing...")
			drone.Land()
		})
	}

	robot := gobot.NewRobot("minidrone",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{drone},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
