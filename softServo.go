package wiringpi

/*
#include <softServo.h>
*/
import "C"

/*

------------------------------------
!!! These functions are disabled !!!
------------------------------------

func SoftServoWrite(pin int, value int) {
    C.softServoWrite(C.int(pin), C.int(value))
}

func SoftServoSetup(p0 int, p1 int, p2 int, p3 int, p4 int, p5 int, p6 int, p7 int) (int, error) {
    retval := int(C.softServoSetup(C.int(p0), C.int(p1), C.int(p2), C.int(p3), C.int(p4), C.int(p5), C.int(p6), C.int(p7)))
    return retval, wiringPiError(retval)
}

*/
