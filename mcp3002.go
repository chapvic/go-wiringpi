package wiringpi

/*
#include <mcp3002.h>
*/
import "C"

// Create a new wiringPi device node for an mcp3002 on the Pi's SPI interface.
func Mcp3002Setup(pinBase int, spiChannel int) (int, error) {
    retval := int(C.mcp3002Setup(C.int(pinBase), C.int(spiChannel)))
    return retval, wiringPiError(retval)
}
