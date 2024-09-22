package wiringpi

/*
#include <mcp3004.h>
*/
import "C"

/*
 * Mcp3004Setup:
 *      Create a new wiringPi device node for an mcp3004 on the Pi's
 *      SPI interface.
 *********************************************************************************
 */

func Mcp3004Setup(pinBase int, spiChannel int) (int, error) {
    retval := int(C.mcp3004Setup(C.int(pinBase), C.int(spiChannel)))
    return retval, wiringPiError(retval)
}
