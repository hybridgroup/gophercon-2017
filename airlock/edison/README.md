# Airlock - Edison

## Installation

```
go get -d -u github.com/hybridgroup/gobot/... && && go install github.com/hybridgroup/gobot/platforms/intel-iot/edison
```

## Running the code
When you run any of these examples, you will compile the code on your computer, move the compiled code onto the Intel Edison, and then execute the code on the Intel Edison itself, not on your own computer.

We have included a shell script to make this process easier. You can run it like this:

```
$ ./runner.sh 1 192.168.1.42
```

For example, to compile the code:

```
$ GOARCH=386 GOOS=linux go build 1.go
```

Then to move the code to the Intel Edison, you can use the `scp` command:

```
$ scp 1 root@<IP of your device>:/home/root/1
```

Lastly, to execute it on your Edison, you can use the `ssh` command:

```
$ ssh -t root@<IP of your device> '1'
```

## Code

### 1.go - Blue LED

Connect the blue LED to D3.

Run the code.

You should see the blue LED blink.

### 2.go - Blue LED, Button

Connect the button to D2.

Run the code.

When you press the button, the blue LED should turn on.

### 3.go - Blue LED, Button, Green LED

Connect the Green LED to D4.

Run the code.

The green LED should light up. When you press the button, the blue LED should turn on, and the green LED should turn off.

### 4.go - Blue LED, Button, Green LED, Cylon.js API

This step has us playing with the Cylon.js API. No additional hardware needs to be connected.

Run the code.

You can now point your web browser to `http://localhost:3000` and try out the [Robeaux](https://github.com/hybridgroup/robeaux) web interface.

### 5.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch

Connect the buzzer to D7, and connect the touch sensor to D8.

Run the code.

When your finger touches the capacitive touch sensor, the buzzer should sound.

### 6.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch, Dial

Connect the rotary dial to A0.

Run the code.

Turning the dial will display the current analog reading on your console.

### 7.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch, Dial, Temperature, Red LED

Connect the temperature sensor to A1, and the red LED to D5

Run the code.

If the temperature exceeds 400 degrees, then the red led will light up.

### 8.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch, Dial, Temperature, Red LED, Sound

Connect the sound sensor to A2.

Run the code.

When a sound is detected, the blue LED will light up, the sound sensor reading will be displayed on your computer's console.

### 9.go - Blue LED, Button, Green LED, Cylon.js API, Buzzer, Touch, Dial, Temperature, Red LED, Sound, Light

Connect the light sensor to A3.

Run the code.

When a light is detected, the blue LED will light up, and the light sensor reading will be displayed on your computer's console.
