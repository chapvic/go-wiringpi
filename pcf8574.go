package wiringpi

/*
#include <pcf8574.h>
*/
import "C"

/*
 * Pcf8574Setup:
 *      Create a new instance of a PCF8574 I2C GPIO interface. We know it
 *      has 8 pins, so all we need to know here is the I2C address and the
 *      user-defined pin base. Default address (A0-A3 low) is 0x20.
 *********************************************************************************
 */

func Pcf8574Setup (pinBase int, i2cAddress int) (int, error) {
    retval := int(C.pcf8574Setup(C.int(pinBase), C.int(i2cAddress)))
    return retval, wiringPiError(retval)
}
