package wiringpi

/*
#include <stdlib.h>
#include <drcSerial.h>
*/
import "C"
import "unsafe"

/*
 * DrcSetupSerial:
 *      Create a new instance of an DRC GPIO interface.
 *      Could be a variable nunber of pins here - we might not know in advance
 *      if it's an ATmega with 14 pins, or something with less or more!
 *********************************************************************************
 */

func DrcSetupSerial(pinBase int, numPins int, device string, baud int) (int, error) {
    _device := C.CString(device)
    defer C.free(unsafe.Pointer(_device))
    retval := int(C.drcSetupSerial(C.int(pinBase), C.int(numPins), (*C.char)(_device), C.int(baud)))
    return retval, wiringPiError(retval)
}