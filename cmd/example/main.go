package main

import (
    "fmt"
    "github.com/chapvic/go-wiringpi"
)

func main() {
    major, minor := wiringpi.Version()
    fmt.Printf("WiringPi Library v%d.%d\n\n", major, minor)

    pins := wiringpi.PiBoardPins()
    if pins < 0 {
        fmt.Println(wiringpi.ERR_PI_BOARD_NOT_DETECTED)
    } else {
        model, rev, maker, mem, over := wiringpi.PiBoardId()
        fmt.Printf("Model       : %s\n", model)
        fmt.Printf("Revision    : %s\n", rev)
        fmt.Printf("Maker       : %s\n", maker)
        fmt.Printf("RAM         : %d\n", mem)
        fmt.Printf("OverVoltage : %d\n", over)
        fmt.Printf("GPIO pins   : %d\n\n", pins)
    }
}
