package wiringpi

/*
#include <sn3218.h>
*/
import "C"

// Create a new wiringPi device node for an sn3218 on the Pi's SPI interface.
func Sn3218Setup (pinBase int) (int, error) {
    retval := int(C.sn3218Setup(C.int(pinBase)))
    return retval, wiringPiError(retval)
}
