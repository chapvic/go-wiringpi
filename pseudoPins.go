package wiringpi

/*
#include <pseudoPins.h>
*/
import "C"

// Create a new wiringPi device node for the pseudoPins driver.
func PseudoPinsSetup (pinBase int) (int, error) {
    retval := int(C.pseudoPinsSetup(C.int(pinBase)))
    return retval, wiringPiError(retval)
}
