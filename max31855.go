package wiringpi

/*
#include <max31855.h>
*/
import "C"

/*
 * Max31855Setup:
 *      Create a new wiringPi device node for an max31855 on the Pi's
 *      SPI interface.
 *********************************************************************************
 */

func Max31855Setup(pinBase int, spiChannel int) (int, error) {
    retval := int(C.max31855Setup(C.int(pinBase), C.int(spiChannel)))
    return retval, wiringPiError(retval)
}
