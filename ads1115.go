package wiringpi

/*
#include <ads1115.h>
*/
import "C"

// Constants for some of the internal functions

const (
    // Gain

    ADS1115_GAIN_6       = 0
    ADS1115_GAIN_4       = 1
    ADS1115_GAIN_2       = 2
    ADS1115_GAIN_1       = 3
    ADS1115_GAIN_HALF    = 4
    ADS1115_GAIN_QUARTER = 5

    // Data rate

    ADS1115_DR_8   = 0
    ADS1115_DR_16  = 1
    ADS1115_DR_32  = 2
    ADS1115_DR_64  = 3
    ADS1115_DR_128 = 4
    ADS1115_DR_250 = 5
    ADS1115_DR_475 = 6
    ADS1115_DR_860 = 7
)

/*
 * Ads1115Setup:
 *      Create a new wiringPi device node for an ads1115 on the Pi's
 *      I2C interface.
 *********************************************************************************
 */

func Ads1115Setup (pinBase int, i2cAddress int) (int, error) {
    retval := int(C.ads1115Setup(C.int(pinBase), C.int(i2cAddress)))
    return retval, wiringPiError(retval)
}
