# Airlock - Edison

## Installation

```
go get -d -u github.com/hybridgroup/gobot/... && && go install github.com/hybridgroup/gobot/platforms/intel-iot/edison
```
## Finding and connecting to the Edison - OSX
- Connect both USB cables to power and establish a serial connection
- Run 
```
command ls /dev/cu.usbserial-*
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

Connect the blue LED to D3.

Run the code.

You should see the blue LED blink.

### step2.go - Blue LED, Button

Connect the button to D2.

Run the code.

When you press the button, the blue LED should turn on.

### step3.go - Blue LED, Button, Green LED

Connect the Green LED to D4.

Run the code.

The green LED should light up. When you press the button, the blue LED should turn on, and the green LED should turn off.

### step4.go - Blue LED, Button, Green LED, Cylon.js API

This step has us playing with the Cylon.js API. No additional hardware needs to be connected.

Run the code.

You can now point your web browser to `http://localhost:3000` and try out the [Robeaux](https://github.com/hybridgroup/robeaux) web interface.

### step5.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch

Connect the buzzer to D7, and connect the touch sensor to D8.

Run the code.

When your finger touches the capacitive touch sensor, the buzzer should sound.

### step6.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch, Dial

Connect the rotary dial to A0.

Run the code.

Turning the dial will display the current analog reading on your console.

### step7.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch, Dial, Temperature, Red LED

Connect the temperature sensor to A1, and the red LED to D5

Run the code.

If the temperature exceeds 400 degrees, then the red led will light up.

### step8.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch, Dial, Temperature, Red LED, Sound

Connect the sound sensor to A2.

Run the code.

When a sound is detected, the blue LED will light up, the sound sensor reading will be displayed on your computer's console.

### step9.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch, Dial, Temperature, Red LED, Sound, Light

Connect the light sensor to A3.

Run the code.

When a light is detected, the blue LED will light up, and the light sensor reading will be displayed on your computer's console.
