package wiringpi

/*
#include <mcp4802.h>
*/
import "C"

// Create a new wiringPi device node for an mcp4802 on the Pi's SPI interface.
func Mcp4802Setup(pinBase int, spiChannel int) (int, error) {
    retval := int(C.mcp4802Setup(C.int(pinBase), C.int(spiChannel)))
    return retval, wiringPiError(retval)
}
