# Gobot Starter Kit 2017

The starter kit for the Gophercon 2017 Community Day!

## Intel TinyTILE

## Sensor Kit

### Sound Sensor

![](../images/starterkit/analog-digital-sound-sensor.png)

module that can be used with analog signal to measure sound threshold and/or with digital to detect if sound was made

### Light Sensor

![](../images/starterkit/photo-resistor-sensor.png)

analog photo luminescence resister module for detecting light

### Magnetic Sensor - Digital

![](../images/starterkit/digital-magnetic-sensor.png)

digital module with an led to detect the presence of a magnetic field

### Magnetic Sensor - Analog

![](../images/starterkit/digital-magnetic-sensor.png)

analog module that is sensative to how near or far a magnet is detected 

### Red/Green LED - Small

![](../images/starterkit/red-green-led.png)

"-" symbol is the ground, middle pin is the posative for green and the end labeled "s" is the posative for the green

### Red/Green LED - Large

![](../images/starterkit/red-green-led.png)

"-" symbol is the ground, middle pin is the posative for green and the end labeled "s" is the posative for the green

### Temperature Sensor

![](../images/starterkit/analog-temperature-sensor.png)

analog temperature sensor 5v module ground is on the "-" symbol 5v is the center pin,and signal is analog on the "s" labeled pin... Need to do proper math to get correct temperature

### Magnetic Sensor - Module

![](../images/starterkit/anaolg-digital-magnetic-feild-sensor.png)

analog/digital magnetic feild module 5v red signal led included. A0 is anlog pin, "G" is ground ", "+" is 5v posative, D0 pin is the digital pin for this module

### Temperature Sensor - Module

![](../images/starterkit/analog-digital-temperature-sensor.png)

analog/digital temperature sensor A0 is anlog pin, "G" is ground ", "+" is 5v posative, D0 pin is the digital pin for this module math needs work to get correct temp

### Touch Sensor - Module

![](../images/starterkit/analog-digital-touch-sensor.png)

analog/digital touch sensor A0 is anlog pin, "G" is ground ", "+" is 5v posative, D0 pin is the digital 

### Magnetic Reed Module

![](../images/starterkit/analog-digital-magnetic-reed-sensor.png)

analog/digital magnetic reed sensor that open or close when exposed to a magnetic feild A0 is anlog pin, "G" is ground ", "+" is 5v posative, D0 pin is the digital red LED signals when magnet is near
 
### RGB LED (flat)

![](../images/starterkit/rgb-flat-led.png)

"-" symbol is ground, "R" symbol is green posative, "G" symbol is red posative, and "B" symbol is blue posative

### Light Cup

![](../images/starterkit/digital-magic-cup-mercury-tilt-sensor.png)

digital gravity powered mercury switch, LED on switch does not appear to work tried 4 different sensors to test "G" is ground ", "+" is 5v posative S is digital signal(works), " L" is 5v led which appears to not work

### Infrared Emission Sensor

![](../images/starterkit/analog-digital-ir-emmission-sensor.png)

analog/digital infrared sensor 5v red signal led included. A0 is anlog pin, "G" is ground ", "+" is 5v posative, D0 pin is the digital pin for this module 

### RGB LED

![](../images/starterkit/rgb-led.png)

this one has the correct labeling yay! "-" symbol is ground, "G" symbol is green posative, "R" symbol is red posative, and "B" symbol is blue posative

### IR Sensor

![](../images/starterkit/digital-ir-reciever-sensor.png)

digital infrared sensor "-" is negative, "s" pin is the digital signal pin and the middle pin is 5v positive

### Magnetic Reed Switch

![](../images/starterkit/digital-magnetic-reed-sensor.png)

digital magnetic reed switch "-" is negative, "s" pin is the digital signal pin and the middle pin is 5v positive

### Mercury Switch
![](../images/starterkit/digital-mercury-tilt-sensor-with-led.png)

digital murcury switch with red LED, "-" is negative, "s" pin is the digital signal pin and the middle pin is 5v positive, tilt the switch to trigger the signal and the LED also turns on

### Hit Sensor - Module
![](../images/starterkit/digital-knock-hit-sensor.png)

digital knock/hit sensor "-" is negative, "s" pin is the digital signal pin and the middle pin is 5v positive, the sensor is not sensitive must hit kinda hard to trigger

### Vibration Sensor
![](../images/starterkit/digital-vibration-sensor.png)

Digital Vibration sensor "-" is negative, "s" pin is the digital signal pin and the middle pin is 5v positive

### Tilt Sensor
![](../images/starterkit/digital-tilt-sensor.png)

digital tilt sensor "-" is negative, "s" pin is the digital signal pin and the middle pin is 5v positive

### Relay
![](../images/starterkit/digital-relay-with-led.png)

digital relay with red led to signal open / close... I couln't get the output to light an led properly may need to test more  "-" is negative, "s" pin is the digital signal pin and the middle pin is 5v positive

### Active Buzzer X2
![](../images/starterkit/active-buzzer.png)

digital buzzer... one buzzer works correctly and the other buzzer "-" symbol == positive

### Push Button
![](../images/starterkit/button.png)

simple push putton two modes either pushed or not pushed "-" is negative, "s" pin is the digital signal pin and the middle pin is 5v positive

### Joystick
![](../images/starterkit/analog-joystick.png)

need to write a proper driver "gnd" is ground, "+5v" is 5v, "vrx" is x axis, "vry" is y axis, sw is presssing the button in

### Optical Break Sensor
![](../images/starterkit/digital-optical-break-sensor.png)

not working atm with current tests

### Obstacle Avoidance Sensor 1
![](../images/starterkit/analog-digital-sound-sensor.png)

digital ir motion sensor "gnd" is ground, "+" is 5v out is digital  signal cable

### Line Finder
![](../images/starterkit/analog-digital-sound-sensor.png)

digital line finder follows black line, need small flathead to adjust the sensativity, "g" is ground, "v+" is 5v positive, "s" is digital signal cable

### Fingerprint Heartbeat Detector

![](../images/starterkit/fingerprint-heartbeat-detector.png)

needs math to work correctly
this is what it looks like in arduino code

```
// Pulse Monitor Test Script

int sensorPin = 0;
double alpha = 0.75;
int period = 100;
double change = 0.0;
double minval = 0.0;
void setup ()
{
  Serial.begin (9600);
}
void loop ()
{
    static double oldValue = 0;
    static double oldChange = 0;
 
    int rawValue = analogRead (sensorPin);
    double value = alpha * oldValue + (1 - alpha) * rawValue;
 
    Serial.print (rawValue);
    Serial.print (",");
    Serial.println (value);
    oldValue = value;
 
    delay (period);
}
```

### Motion Sensor
![](../images/starterkit/digital-ir-motion-sensor.png)

appears to be an ir motion detector

### LASER
![](../images/starterkit/digital-laser-red.png)

digital red laser "-" is ground, middle wire is 5v positive, "s" is digital signal cable

### LED
![](../images/starterkit/led-no-resistor.png)

LED appears not to work may need resistor

### 1-Wire Temperature Sensor
![](../images/starterkit/digital-temp-sensor.png)

Not yet supported... help wanted! Ask us

### 1-Wire Temperature/Humidity Sensor
![](../images/starterkit/humidity-temperature-sensor.png)

Not yet supported... help wanted! Ask us

### Rotary Encoder
![](../images/starterkit/rotary-encoder.png)

not yet supported... help wanted! Ask us
looking for pulse encoding support