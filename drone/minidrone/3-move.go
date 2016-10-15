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
		drone.On(drone.Event("battery"), func(data interface{}) {
			fmt.Printf("battery: %d\n", data)
		})

		drone.On(drone.Event("status"), func(data interface{}) {
			fmt.Printf("status: %d\n", data)
		})

		drone.On(drone.Event("flying"), func(data interface{}) {
			fmt.Println("flying!")
			gobot.After(5*time.Second, func() {
				fmt.Println("forwards...")
				drone.Forward(10)
			})
			gobot.After(10*time.Second, func() {
				fmt.Println("backwards...")
				drone.Backward(10)
			})
			gobot.After(15*time.Second, func() {
				fmt.Println("right...")
				drone.Right(10)
			})
			gobot.After(20*time.Second, func() {
				fmt.Println("left...")
				drone.Left(10)
			})
			gobot.After(25*time.Second, func() {
				fmt.Println("landing...")
				drone.Land()
			})
		})

		drone.On(drone.Event("landed"), func(data interface{}) {
			fmt.Println("landed.")
		})

		drone.TakeOff()
	}

	robot := gobot.NewRobot("minidrone",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{drone},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
