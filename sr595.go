package wiringpi

/*
#include <sr595.h>
*/
import "C"

// Create a new instance of a 74x595 shift register GPIO expander.
func Sr595Setup (pinBase int, numPins int, dataPin int, clockPin int, latchPin int) (int, error) {
    retval := int(C.sr595Setup(C.int(pinBase), C.int(numPins), C.int(dataPin), C.int(clockPin), C.int(latchPin)))
    return retval, wiringPiError(retval)
}
