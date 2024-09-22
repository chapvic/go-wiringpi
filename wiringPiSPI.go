package wiringpi

/*
#include <stdlib.h>
#include <wiringPiSPI.h>
*/
import "C"
import "unsafe"

func SPISetup(channel int, speed int) (int, error) {
    retval := int(C.wiringPiSPISetup(C.int(channel), C.int(speed)))
    return retval, wiringPiError(retval)
}

func SPISetupMode(channel int, speed int, mode int) (int, error) {
    retval := int(C.wiringPiSPISetupMode(C.int(channel), C.int(speed), C.int(mode)))
    return retval, wiringPiError(retval)
}

func SPIClose(channel int) (int, error) {
    retval := int(C.wiringPiSPIClose(C.int(channel)))   //Interface 3.5
    return retval, wiringPiError(retval)
}

func SPIGetFd(channel int) (int, error) {
    retval := int(C.wiringPiSPIGetFd(C.int(channel)))
    return retval, wiringPiError(retval)
}

func SPIDataRW(channel int, data string) (int, error) {
    clen := len(data)
    cstr := C.CString(data)
    defer C.free(unsafe.Pointer(cstr))
    retval := int(C.wiringPiSPIDataRW(C.int(channel), (*C.uchar)(cstr), C.int(clen)))
    return retval, wiringPiError(retval)
}

// Interface 3.5

func SPIxGetFd(number int, channel int) (int, error) {
    retval := int(C.wiringPiSPIxGetFd(C.int(number), C.int(channel)))
    return retval, wiringPiError(retval)
}

func SPIxDataRW(number int, channel int, data string) (int, error) {
    clen := len(data)
    cstr := C.CString(data)
    defer C.free(unsafe.Pointer(cstr))
    retval := int(C.wiringPiSPIxDataRW(C.int(number), C.int(channel), (*C.uchar)(cstr), C.int(clen)))
    return retval, wiringPiError(retval)
}

func SPIxSetupMode(number int, channel int, speed int, mode int) (int, error) {
    retval := int(C.wiringPiSPIxSetupMode(C.int(number), C.int(channel), C.int(speed), C.int(mode)))
    return retval, wiringPiError(retval)
}

/*
WARNING: This function is declared, but not implemented!

func SPIxSetup(number int, channel int, speed int) (int, error) {
    retval := int(C.wiringPiSPIxSetup(C.int(number), C.int(channel), C.int(speed)))
    return retval, wiringPiError(retval)
}
*/

func SPIxClose(number int, channel int) (int, error) {
    retval := int(C.wiringPiSPIxClose(C.int(number), C.int(channel)))
    return retval, wiringPiError(retval)
}
