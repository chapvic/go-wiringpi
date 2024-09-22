package wiringpi

/*
#include <wiringPiI2C.h>
*/
import "C"

func I2CRead(fd int) (int, error) {
    retval := int(C.wiringPiI2CRead(C.int(fd)))
    return retval, wiringPiError(retval)
}

func I2CReadReg8(fd int, reg int) (int, error) {
    retval := int(C.wiringPiI2CReadReg8(C.int(fd), C.int(reg)))
    return retval, wiringPiError(retval)
}

func I2CReadReg16(fd int, reg int) (int, error) {
    retval := int(C.wiringPiI2CReadReg16(C.int(fd), C.int(reg)))
    return retval, wiringPiError(retval)
}

func I2CReadBlockData(fd int, reg int, values *uint8, size uint8) (int, error) {
    retval := int(C.wiringPiI2CReadBlockData(C.int(fd), C.int(reg), (*C.uint8_t)(values), C.uint8_t(size)))        // Interface 3.3
    return retval, wiringPiError(retval)
}

func I2CRawRead(fd int, values *uint8, size uint8) (int, error) {
    retval := int(C.wiringPiI2CRawRead(C.int(fd), (*C.uint8_t)(values), C.uint8_t(size)))                          // Interface 3.3
    return retval, wiringPiError(retval)
}

func I2CWrite(fd int, data int) (int, error) {
    retval := int(C.wiringPiI2CWrite(C.int(fd), C.int(data)))
    return retval, wiringPiError(retval)
}

func I2CWriteReg8(fd int, reg int, data int) (int, error) {
    retval := int(C.wiringPiI2CWriteReg8(C.int(fd), C.int(reg), C.int(data)))
    return retval, wiringPiError(retval)
}

func I2CWriteReg16(fd int, reg int, data int) (int, error) {
    retval := int(C.wiringPiI2CWriteReg16(C.int(fd), C.int(reg), C.int(data)))
    return retval, wiringPiError(retval)
}

func I2CWriteBlockData(fd int, reg int, values *uint8, size uint8) (int, error) {
    retval := int(C.wiringPiI2CWriteBlockData(C.int(fd), C.int(reg), (*C.uint8_t)(values), C.uint8_t(size)))       // Interface 3.3
    return retval, wiringPiError(retval)
}

func I2CRawWrite(fd int, values *uint8, size uint8) (int, error) {
    retval := int(C.wiringPiI2CRawWrite(C.int(fd), (*C.uint8_t)(values), C.uint8_t(size)))                         // Interface 3.3
    return retval, wiringPiError(retval)
}

func I2CSetupInterface(device *byte, devId int) (int, error) {
    retval := int(C.wiringPiI2CSetupInterface((*C.char)(device), C.int(devId)))
    return retval, wiringPiError(retval)
}

func I2CSetup(devId int) (int, error) {
    retval := int(C.wiringPiI2CSetup(C.int(devId)))
    return retval, wiringPiError(retval)
}
