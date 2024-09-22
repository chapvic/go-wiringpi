package wiringpi

/*
#include <softPwm.h>
*/
import "C"

/*
 * SoftPwmCreate:
 *      Create a new softPWM thread.
 *********************************************************************************
 */

func SoftPwmCreate(pin int, value int, datarange int) (int, error) {
    retval := int(C.softPwmCreate (C.int(pin), C.int(value), C.int(datarange)))
    return retval, wiringPiError(retval)
}

/*
 * SoftPwmWrite:
 *      Write a PWM value to the given pin
 *********************************************************************************
 */

func SoftPwmWrite(pin int, value int) {
    C.softPwmWrite(C.int(pin), C.int(value))
}

/*
 * SoftPwmStop:
 *      Stop an existing softPWM thread
 *********************************************************************************
 */

func SoftPwmStop(pin int) {
    C.softPwmStop(C.int(pin))
}
