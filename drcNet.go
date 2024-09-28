package wiringpi

/*
#include <stdlib.h>
#include <drcNet.h>
*/
import "C"
import "unsafe"

// Create a new instance of an DRC GPIO interface.
// Could be a variable nunber of pins here - we might not know in advance.
func DrcSetupNet(pinBase int, numPins int, ipAddress string, port string, password string) (int, error) {
    _ipAddress := C.CString(ipAddress)
    defer C.free(unsafe.Pointer(_ipAddress))
    _port := C.CString(port)
    defer C.free(unsafe.Pointer(_port))
    _password := C.CString(password)
    defer C.free(unsafe.Pointer(_password))
    retval := int(C.drcSetupNet(C.int(pinBase), C.int(numPins), (*C.char)(_ipAddress), (*C.char)(_port), (*C.char)(_password)))
    return retval, wiringPiError(retval)
}
