package wiringpi

/*
#include <bmp180.h>
*/
import "C"

/*
 * Bmp180Setup:
 *      Create a new instance of a PCF8591 I2C GPIO interface. We know it
 *      has 4 pins, (4 analog inputs and 1 analog output which we'll shadow
 *      input 0) so all we need to know here is the I2C address and the
 *      user-defined pin base.
 *********************************************************************************
 */

func Bmp180Setup(pinBase int) (int, error) {
    retval := int(C.bmp180Setup(C.int(pinBase)))
    return retval, wiringPiError(retval)
}
