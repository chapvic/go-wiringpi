package wiringpi

/*
#include <mcp23s17.h>
*/
import "C"

/*
 * Mcp23s17Setup:
 *      Create a new instance of an MCP23s17 SPI GPIO interface. We know it
 *      has 16 pins, so all we need to know here is the SPI address and the
 *      user-defined pin base.
 *********************************************************************************
 */

func Mcp23s17Setup(pinBase int, spiPort int, devId int) (int, error) {
    retval := int(C.mcp23s17Setup(C.int(pinBase), C.int(spiPort), C.int(devId)))
    return retval, wiringPiError(retval)
}
