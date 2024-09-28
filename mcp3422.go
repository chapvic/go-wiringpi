package wiringpi

/*
#include <mcp3422.h>
*/
import "C"

const (
    // Sample Rate

    MCP3422_SR_240      = 0
    MCP3422_SR_60       = 1
    MCP3422_SR_15       = 2
    MCP3422_SR_3_75     = 3

    // Gain

    MCP3422_GAIN_1      = 0
    MCP3422_GAIN_2      = 1
    MCP3422_GAIN_4      = 2
    MCP3422_GAIN_8      = 3
)

// Create a new wiringPi device node for the mcp3422
func Mcp3422Setup(pinBase int, i2cAddress int, sampleRate int, gain int) (int, error) {
    retval := int(C.mcp3422Setup(C.int(pinBase), C.int(i2cAddress), C.int(sampleRate), C.int(gain)))
    return retval, wiringPiError(retval)
}
