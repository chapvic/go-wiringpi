package main

import (
    "fmt"
    "github.com/chapvic/go-wiringpi"
)

func main() {
    major, minor := wiringpi.Version()
    fmt.Printf("WiringPi Library v%d.%d\n\n", major, minor)

    model, rev, maker, mem, over := wiringpi.PiBoardId()
    fmt.Printf("Model       : %s\n", model)
    fmt.Printf("Revision    : %s\n", rev)
    fmt.Printf("Maker       : %s\n", maker)
    fmt.Printf("RAM         : %d\n", mem)
    fmt.Printf("OverVoltage : %d\n", over)
}
