package wiringpi

/*
#include <pcf8591.h>
*/
import "C"

/*
 * Pcf8591Setup:
 *      Create a new instance of a PCF8591 I2C GPIO interface. We know it
 *      has 4 pins, (4 analog inputs and 1 analog output which we'll shadow
 *      input 0) so all we need to know here is the I2C address and the
 *      user-defined pin base.
 *********************************************************************************
 */

func Pcf8591Setup (pinBase int, i2cAddress int) (int, error) {
    retval := int(C.pcf8591Setup(C.int(pinBase), C.int(i2cAddress)))
    return retval, wiringPiError(retval)
}
