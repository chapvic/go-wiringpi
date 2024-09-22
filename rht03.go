package wiringpi

/*
#include <rht03.h>
*/
import "C"

/*
 * Rht03Setup:
 *      Create a new instance of an RHT03 temperature sensor.
 *********************************************************************************
 */

func Rht03Setup (pinBase int, devicePin int) (int, error) {
    retval := int(C.rht03Setup(C.int(pinBase), C.int(devicePin)))
    return retval, wiringPiError(retval)
}
