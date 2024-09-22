package wiringpi

/*
#include <htu21d.h>
*/
import "C"

/*
 * Htu21dSetup:
 *      Create a new instance of a HTU21D I2C GPIO interface.
 *      This chip has a fixed I2C address, so we are not providing any
 *      allowance to change this.
 *********************************************************************************
 */

func Htu21dSetup(pinBase int) (int, error) {
    retval := int(C.htu21dSetup(C.int(pinBase)))
    return retval, wiringPiError(retval)
}
