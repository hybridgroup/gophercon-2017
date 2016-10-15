# Rover - Sphero Ollie

## Installation

```
go get -d -u github.com/hybridgroup/gobot/... && go install github.com/hybridgroup/gobot/platforms/ble
```

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the drone  using the Bluetooth Low Energy (LE) interface.

To compile/run the code, substitute the name of your Ollie or BB-8 as needed.

### OS X

```
$ GODEBUG=cgocheck=0 go run 1-color.go BB-128E
```

### Linux

```
$ go run 1-color.go BB-128E
```

## Code

### 1-color.go

### 2-roll.go

## License

Copyright (c) 2015-2016 The Hybrid Group. Licensed under the MIT license.
