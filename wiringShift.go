package wiringpi

/*
#include <wiringShift.h>
*/
import "C"

// Shift orders
const (
    LSBFIRST = 0
    MSBFIRST = 1
)

func shiftIn(dPin int, cPin int, order int) int {
    return int(C.shiftIn(C.uint8_t(dPin), C.uint8_t(cPin), C.uint8_t(order)))
}

func shiftOut(dPin int, cPin int, order int, val int) {
    C.shiftOut(C.uint8_t(dPin), C.uint8_t(cPin), C.uint8_t(order), C.uint8_t(val))
}
