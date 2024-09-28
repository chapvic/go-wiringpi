package wiringpi

/*
#include <softTone.h>
*/
import "C"

// Create a new tone thread.
func SoftToneCreate(pin int) (int, error) {
    retval := int(C.softToneCreate(C.int(pin)))
    return retval, wiringPiError(retval)
}

// Stop an existing softTone thread.
func SoftToneStop(pin int) {
    C.softToneStop(C.int(pin))
}

// Write a frequency value to the given pin.
func SoftToneWrite(pin int, freq int) {
    C.softToneWrite(C.int(pin), C.int(freq))
}
