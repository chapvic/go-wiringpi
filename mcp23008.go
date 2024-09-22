package wiringpi

/*
#include <mcp23008.h>
*/
import "C"

/*
 * Mcp23008Setup:
 *      Create a new instance of an MCP23008 I2C GPIO interface. We know it
 *      has 8 pins, so all we need to know here is the I2C address and the
 *      user-defined pin base.
 *********************************************************************************
 */

func Mcp23008Setup (pinBase int, i2cAddress int) (int, error) {
    retval := int(C.mcp23008Setup(C.int(pinBase), C.int(i2cAddress)))
    return retval, wiringPiError(retval)
}
