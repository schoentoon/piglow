piglow
======

[![GoDoc](https://godoc.org/github.com/schoentoon/piglow?status.svg)](https://godoc.org/github.com/schoentoon/piglow)

A PiGlow library written in Go


You can cross compile this library/example using the following command, assuming you have the cross compilers for golang installed on your system.

```
GOARCH=arm GOARM=6 go build
```

To include this library into your own project simply use the following statement.

```
import "github.com/schoentoon/piglow"
```

Usage is pretty straight forward, use PiGlow(led, intensity) to set individual LEDs and call ShutDown() to turn them all off. An example can be found in the source tree under main.go

This project is mostly based on https://github.com/pimoroni/piglow/blob/master/examples/piglow-example.py
