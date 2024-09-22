package wiringpi

/*
#include <mcp23016.h>
*/
import "C"

/*
 * Mcp23016Setup:
 *      Create a new instance of an MCP23016 I2C GPIO interface. We know it
 *      has 16 pins, so all we need to know here is the I2C address and the
 *      user-defined pin base.
 *********************************************************************************
 */

func Mcp23016Setup (pinBase int, i2cAddress int) (int, error) {
    retval := int(C.mcp23016Setup(C.int(pinBase), C.int(i2cAddress)))
    return retval, wiringPiError(retval)
}
