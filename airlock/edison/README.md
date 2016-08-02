# Airlock - Edison

## Installation

```
go get -d -u github.com/hybridgroup/gobot/... && go install github.com/hybridgroup/gobot/platforms/intel-iot/edison
go get -d -u github.com/hybridgroup/gobot-workshop-2016
```
## Setting up your board
Follow the directions at the Edison site  
`https://software.intel.com/en-us/node/628232`

or just download the setup from and follow the setup prompts `https://software.intel.com/iot/hardware/edison/downloads`

## Finding and connecting to the Edison - OSX
- Connect both USB cables to power and establish a serial connection
- Run
```
ls /dev/cu.usbserial-*
```
- Copy the cu.usbserial-XXXXXX that is returned
- Run
```
screen /dev/xx.usbserial-XXXXXXXX 115200 –L
```
- Press enter twice to login
- Run ifconfig command and look for wlan0 entry

## Running the code
When you run any of these examples, you will compile the code on your computer, move the compiled code onto the Intel Edison, and then execute the code on the Intel Edison itself, not on your own computer.

We have included a shell script to make this process easier. You can run it like this:

```
$ ./runner.sh step1 192.168.1.42
```
Note: You'll use the IP Address you get during the setup process

The `runner.sh` script performs the following steps for you:

To compile the code:

```
$ GOARCH=386 GOOS=linux go build step1.go
```

Then to move the code to the Intel Edison, it uses the `scp` command:

```
$ scp step1 root@<IP of your device>:/home/root/step1
```

Lastly, to execute it on your Edison, it uses the `ssh` command:

```
$ ssh -t root@<IP of your device> ./step1
```

## Code

### step1.go - Blue LED

![Gobot](https://raw.githubusercontent.com/hybridgroup/gobot-workshop-2016/master/images/edison/step1.jpg)

Connect the blue LED to D3.

Run the code.

You should see the blue LED blink.

### step2.go - Blue LED, Button

![Gobot](https://raw.githubusercontent.com/hybridgroup/gobot-workshop-2016/master/images/edison/step2.jpg)

Connect the button to D2.

Run the code.

When you press the button, the blue LED should turn on.

### step3.go - Blue LED, Button, Green LED

![Gobot](https://raw.githubusercontent.com/hybridgroup/gobot-workshop-2016/master/images/edison/step3.jpg)

Connect the Green LED to D4.

Run the code.

The green LED should light up. When you press the button, the blue LED should turn on, and the green LED should turn off.

### step4.go - Blue LED, Button, Green LED, Gobot API

This step has us playing with the Gobot API. No additional hardware needs to be connected.

Run the code.

You can now point your web browser to `http://<IP of your device>:3000` and try out the [Robeaux](https://github.com/hybridgroup/robeaux) web interface.

### step5.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch

![Gobot](https://raw.githubusercontent.com/hybridgroup/gobot-workshop-2016/master/images/edison/step5.jpg)

Connect the buzzer to D6, and connect the touch sensor to D8.

Run the code.

When your finger touches the capacitive touch sensor, the buzzer should sound.

### step6.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial

![Gobot](https://raw.githubusercontent.com/hybridgroup/gobot-workshop-2016/master/images/edison/step6.jpg)

Connect the rotary dial to A0.

Run the code.

Turning the dial will display the current analog reading on your console.

### step7.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial, Temperature, Red LED

![Gobot](https://raw.githubusercontent.com/hybridgroup/gobot-workshop-2016/master/images/edison/step7.jpg)

Connect the temperature sensor to A1, and the red LED to D5

Run the code.

By default, if the temperature exceeds 15 degrees, then the red LED will light up.
In case the room is warmer than 15 degrees, change the default temperature in step7.go.

In order to increase the temperature of the sensor, hold it between your fingers and wait for the LED to light up.
To turn the LED off, let go of the temperature sensor (note: the temperature will drop much slower than it increased).

### step8.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial, Temperature, Red LED, Sound

![Gobot](https://raw.githubusercontent.com/hybridgroup/gobot-workshop-2016/master/images/edison/step8.jpg)

Connect the sound sensor to A2.

Run the code.

When a sound is detected, the blue LED will light up, the sound sensor reading will be displayed on your computer's console.

### step9.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial, Temperature, Red LED, Sound, Light

![Gobot](https://raw.githubusercontent.com/hybridgroup/gobot-workshop-2016/master/images/edison/step9.jpg)

Connect the light sensor to A3.

Run the code.

When a light is detected, the blue LED will light up, and the light sensor reading will be displayed on your computer's console.

### step10.go - Blue LED, Button, Green LED, Gobot API, Buzzer, Touch, Dial, Temperature, Red LED, Sound, Light, LCD

![Gobot](https://raw.githubusercontent.com/hybridgroup/gobot-workshop-2016/master/images/edison/step10.jpg)

Connect the LCD to any of the I2C ports on the Grove shield.

Run the code.

The LCD display will show informative messages, and also change the backlight color to match the color of whichever of the 3 LEDs is lit.
