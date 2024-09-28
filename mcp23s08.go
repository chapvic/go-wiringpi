package wiringpi

/*
#include <mcp23s08.h>
*/
import "C"

// Create a new instance of an MCP23s08 SPI GPIO interface.
// We know it has 8 pins, so all we need to know here is the SPI address and the user-defined pin base.
func Mcp23s08Setup(pinBase int, spiPort int, devId int) (int, error) {
    retval := int(C.mcp23s08Setup(C.int(pinBase), C.int(spiPort), C.int(devId)))
    return retval, wiringPiError(retval)
}
