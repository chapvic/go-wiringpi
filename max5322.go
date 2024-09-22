package wiringpi

/*
#include <max5322.h>
*/
import "C"

/*
 * max5322Setup:
 *      Create a new wiringPi device node for an max5322 on the Pi's
 *      SPI interface.
 *********************************************************************************
 */

func Max5322Setup(pinBase int, spiChannel int) (int, error) {
    retval := int(C.max5322Setup(C.int(pinBase), C.int(spiChannel)))
    return retval, wiringPiError(retval)
}
