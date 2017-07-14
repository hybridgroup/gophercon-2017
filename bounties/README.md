# Bounties 2017

Like last year, we have some equipment "bounties". If you can contribute back to Gobot by adding software support you can take home one of the following awesome new devices:

   - GoPiGo3 Raspberry Pi powered Robot
   - Holy Stone HS220 Drone
   - Ultrasonic Rangefinder
   - Stepper Motor Controller

## GoPiGo3

The GoPiGo3 is a Raspberry Pi powered robot kit. It uses a separate GoPiGo3 daughterboard to communicate between the Raspberry Pi and the various motors/sensors/etc.

Some initial work on add the previous version of the GoPiGo was done in this PR, as a possible starting out point for this bounty:

https://github.com/hybridgroup/gobot/pull/338

Your mission is to add a Gobot adaptor/driver for the GoPiGo3 so we can run Gobot on the robot itself.

## Holy Stone HS220 Drone

The Holy Stone HS110W/HS220 is a small battery operated quadcopter that has a connected "brain" and camera. The controller provides a WiFi access point and some kind of UDP based API. The API is not oficially documented, however information on one person's reverse engineering of its API using Node.js can be found here:

https://github.com/lancecaraccioli/holystone-hs110w

Your mission is to add a Gobot adaptor/driver for the HS110W/HS220 so we can run Gobot as a ground control station to control the drone.

## Ultrasonic Rangefinder

## Stepper Motor Controller
