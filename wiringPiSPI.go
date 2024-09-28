package wiringpi

/*
#include <stdlib.h>
#include <wiringPiSPI.h>
*/
import "C"
import "unsafe"

// Open the SPI device, and set it up, etc. in the default MODE 0.
func SPISetup(channel int, speed int) (int, error) {
    retval := int(C.wiringPiSPISetup(C.int(channel), C.int(speed)))
    return retval, wiringPiError(retval)
}

// Open the SPI device, and set it up, with the mode, etc.
func SPISetupMode(channel int, speed int, mode int) (int, error) {
    retval := int(C.wiringPiSPISetupMode(C.int(channel), C.int(speed), C.int(mode)))
    return retval, wiringPiError(retval)
}

// Close the SPI device.
func SPIClose(channel int) (int, error) {
    retval := int(C.wiringPiSPIClose(C.int(channel)))   //Interface 3.5
    return retval, wiringPiError(retval)
}

// Return the file-descriptor for the given channel
func SPIGetFd(channel int) (int, error) {
    retval := int(C.wiringPiSPIGetFd(C.int(channel)))
    return retval, wiringPiError(retval)
}

// Write and Read a block of data over the SPI bus.
// Note the data ia being read into the transmit buffer, so will overwrite it!
// This is also a full-duplex operation.
func SPIDataRW(channel int, data string) (int, error) {
    clen := len(data)
    cstr := C.CString(data)
    defer C.free(unsafe.Pointer(cstr))
    retval := int(C.wiringPiSPIDataRW(C.int(channel), (*C.uchar)(cstr), C.int(clen)))
    return retval, wiringPiError(retval)
}


// Interface 3.5

// Return the file-descriptor for the given channel
func SPIxGetFd(number int, channel int) (int, error) {
    retval := int(C.wiringPiSPIxGetFd(C.int(number), C.int(channel)))
    return retval, wiringPiError(retval)
}

// Write and Read a block of data over the SPI bus.
// Note the data ia being read into the transmit buffer, so will overwrite it!
// This is also a full-duplex operation.
func SPIxDataRW(number int, channel int, data string) (int, error) {
    clen := len(data)
    cstr := C.CString(data)
    defer C.free(unsafe.Pointer(cstr))
    retval := int(C.wiringPiSPIxDataRW(C.int(number), C.int(channel), (*C.uchar)(cstr), C.int(clen)))
    return retval, wiringPiError(retval)
}

// Open the SPI device, and set it up, with the mode, etc.
func SPIxSetupMode(number int, channel int, speed int, mode int) (int, error) {
    retval := int(C.wiringPiSPIxSetupMode(C.int(number), C.int(channel), C.int(speed), C.int(mode)))
    return retval, wiringPiError(retval)
}

/*
This function is declared, but not implemented!

func SPIxSetup(number int, channel int, speed int) (int, error) {
    retval := int(C.wiringPiSPIxSetup(C.int(number), C.int(channel), C.int(speed)))
    return retval, wiringPiError(retval)
}
*/

// Close the SPI device.
func SPIxClose(number int, channel int) (int, error) {
    retval := int(C.wiringPiSPIxClose(C.int(number), C.int(channel)))
    return retval, wiringPiError(retval)
}
