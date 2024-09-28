package wiringpi

/*
#include <stdlib.h>
#include <ds18b20.h>
*/
import "C"
import "unsafe"

// Create a new instance of a DS18B20 temperature sensor.
func Ds18b20Setup(pinBase int, serialNum string) (int, error) {
    _serialNum := C.CString(serialNum)
    defer C.free(unsafe.Pointer(_serialNum))
    retval := int(C.ds18b20Setup(C.int(pinBase), (*C.char)(_serialNum)))
    return retval, wiringPiError(retval)
}
