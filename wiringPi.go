package wiringpi

/*
#cgo LDFLAGS: -lwiringPi
#include <string.h>
#include <wiringPi.h>
*/
import "C"
import (
    "errors"
    "sync"
)

const (

    // wiringPi modes

    WPI_MODE_PINS             = 0
    WPI_MODE_GPIO             = 1
    WPI_MODE_GPIO_SYS         = 2  // deprecated since 3.2
    WPI_MODE_PHYS             = 3
    WPI_MODE_PIFACE           = 4
    WPI_MODE_GPIO_DEVICE_BCM  = 5  // BCM pin numbers like WPI_MODE_GPIO
    WPI_MODE_GPIO_DEVICE_WPI  = 6  // WiringPi pin numbers like WPI_MODE_PINS
    WPI_MODE_GPIO_DEVICE_PHYS = 7  // Physic pin numbers like WPI_MODE_PHYS
    WPI_MODE_UNINITIALISED    = -1

    // Pin modes

    INPUT            = 0
    OUTPUT           = 1
    PWM_OUTPUT       = 2
    PWM_MS_OUTPUT    = 8
    PWM_BAL_OUTPUT   = 9
    GPIO_CLOCK       = 3
    SOFT_PWM_OUTPUT  = 4
    SOFT_TONE_OUTPUT = 5
    PWM_TONE_OUTPUT  = 6
    PM_OFF           = 7   // to input / release line

    LOW              = 0
    HIGH             = 1

    // Pull up/down/none

    PUD_OFF          = 0
    PUD_DOWN         = 1
    PUD_UP           = 2

    // PWM

    PWM_MODE_MS      = 0
    PWM_MODE_BAL     = 1

    // Interrupt levels

    INT_EDGE_SETUP    = 0
    INT_EDGE_FALLING  = 1
    INT_EDGE_RISING   = 2
    INT_EDGE_BOTH     = 3

    // Pi model types and version numbers
    //    Intended for the GPIO program Use at your own risk.
    // https://www.raspberrypi.com/documentation/computers/raspberry-pi.html#new-style-revision-codes

    PI_MODEL_A        = 0
    PI_MODEL_B        = 1
    PI_MODEL_AP       = 2
    PI_MODEL_BP       = 3
    PI_MODEL_2        = 4
    PI_ALPHA          = 5
    PI_MODEL_CM       = 6
    PI_MODEL_07       = 7
    PI_MODEL_3B       = 8
    PI_MODEL_ZERO     = 9
    PI_MODEL_CM3      = 10
    PI_MODEL_ZERO_W   = 12
    PI_MODEL_3BP      = 13
    PI_MODEL_3AP      = 14
    PI_MODEL_CM3P     = 16
    PI_MODEL_4B       = 17
    PI_MODEL_ZERO_2W  = 18
    PI_MODEL_400      = 19
    PI_MODEL_CM4      = 20
    PI_MODEL_CM4S     = 21
    PI_MODEL_5        = 23

    PI_VERSION_1      = 0
    PI_VERSION_1_1    = 1
    PI_VERSION_1_2    = 2
    PI_VERSION_2      = 3

    PI_MAKER_SONY     = 0
    PI_MAKER_EGOMAN   = 1
    PI_MAKER_EMBEST   = 2
    PI_MAKER_UNKNOWN  = 3

    GPIO_LAYOUT_PI1_REV1 = 1   //Pi 1 A/B Revision 1, 1.1, CM
    GPIO_LAYOUT_DEFAULT  = 2


    // Matching BCM pins to WiringPi pins

    GPIO_0  = 30
    GPIO_1  = 31
    GPIO_2  = 8
    GPIO_3  = 9
    GPIO_4  = 7
    GPIO_5  = 21
    GPIO_6  = 22
    GPIO_7  = 11
    GPIO_8  = 10
    GPIO_9  = 13
    GPIO_10 = 12
    GPIO_11 = 14
    GPIO_12 = 26
    GPIO_13 = 23
    GPIO_14 = 15
    GPIO_15 = 16
    GPIO_16 = 27
    GPIO_17 = 0
    GPIO_18 = 1
    GPIO_19 = 24
    GPIO_20 = 28
    GPIO_21 = 29
    GPIO_22 = 3
    GPIO_23 = 4
    GPIO_24 = 5
    GPIO_25 = 6
    GPIO_26 = 25
    GPIO_27 = 2
)

/*
 $ gpio readall
 +-----+-----+---------+------+---+---Pi 5---+---+------+---------+-----+-----+
 | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
 +-----+-----+---------+------+---+----++----+---+------+---------+-----+-----+
 |     |     |    3.3v |      |   |  1 || 2  |   |      | 5v      |     |     |
 |   2 |   8 |   SDA.1 | ALT3 | 1 |  3 || 4  |   |      | 5v      |     |     |
 |   3 |   9 |   SCL.1 | ALT3 | 1 |  5 || 6  |   |      | 0v      |     |     |
 |   4 |   7 | GPIO. 7 |   IN | 1 |  7 || 8  | 0 |  -   | TxD     | 15  | 14  |
 |     |     |      0v |      |   |  9 || 10 | 0 |  -   | RxD     | 16  | 15  |
 |  17 |   0 | GPIO. 0 |   -  | 0 | 11 || 12 | 0 |  -   | GPIO. 1 | 1   | 18  |
 |  27 |   2 | GPIO. 2 |   -  | 0 | 13 || 14 |   |      | 0v      |     |     |
 |  22 |   3 | GPIO. 3 |   -  | 0 | 15 || 16 | 0 | IN   | GPIO. 4 | 4   | 23  |
 |     |     |    3.3v |      |   | 17 || 18 | 0 |  -   | GPIO. 5 | 5   | 24  |
 |  10 |  12 |    MOSI | ALT0 | 0 | 19 || 20 |   |      | 0v      |     |     |
 |   9 |  13 |    MISO | ALT0 | 0 | 21 || 22 | 0 |  -   | GPIO. 6 | 6   | 25  |
 |  11 |  14 |    SCLK | ALT0 | 0 | 23 || 24 | 1 | OUT  | CE0     | 10  | 8   |
 |     |     |      0v |      |   | 25 || 26 | 1 | IN   | CE1     | 11  | 7   |
 |   0 |  30 |   SDA.0 |   IN | 1 | 27 || 28 | 1 | IN   | SCL.0   | 31  | 1   |
 |   5 |  21 | GPIO.21 |   -  | 0 | 29 || 30 |   |      | 0v      |     |     |
 |   6 |  22 | GPIO.22 |   -  | 0 | 31 || 32 | 0 |  -   | GPIO.26 | 26  | 12  |
 |  13 |  23 | GPIO.23 |   -  | 0 | 33 || 34 |   |      | 0v      |     |     |
 |  19 |  24 | GPIO.24 |   -  | 0 | 35 || 36 | 0 |  -   | GPIO.27 | 27  | 16  |
 |  26 |  25 | GPIO.25 |   -  | 0 | 37 || 38 | 0 |  -   | GPIO.28 | 28  | 20  |
 |     |     |      0v |      |   | 39 || 40 | 0 |  -   | GPIO.29 | 29  | 21  |
 +-----+-----+---------+------+---+----++----+---+------+---------+-----+-----+
 | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
 +-----+-----+---------+------+---+---Pi 5---+---+------+---------+-----+-----+

*/

// WPIPinAlt
const (
    WPI_ALT_UNKNOWN = -1
    WPI_ALT_INPUT = 0
    WPI_ALT_OUTPUT = 1
    WPI_ALT5 = 2
    WPI_ALT4 = 3
    WPI_ALT0 = 4
    WPI_ALT1 = 5
    WPI_ALT2 = 6
    WPI_ALT3 = 7
    WPI_ALT6 = 8
    WPI_ALT7 = 9
    WPI_ALT8 = 10
    WPI_ALT9 = 11
    WPI_NONE = 0x1F   // Pi5 default
)

/*
 * Checking the values returned by functions
 * If the return value is < 0, an error message is returned
 */

var _mutex = new(sync.Mutex)

func wiringPiError(retval int) error {
    if retval >= 0 {
        return nil
    }
    _mutex.Lock()
    defer _mutex.Unlock()
    return errors.New(C.GoString(C.strerror(C.int(retval))))
}


/*
 *
 * Wrapper functoins for WiringPi interactions
 *
 */

// Core wiringPi functions

func Version() (int, int) {
    var major, minor C.int
    C.wiringPiVersion(&major, &minor)
    return int(major), int(minor)
}

func GlobalMemoryAccess() int {
    return int(C.wiringPiGlobalMemoryAccess())  // Interface V3.3
}

func UserLevelAccess() int {
    return int(C.wiringPiUserLevelAccess())
}

func Setup() error {
    retval := int(C.wiringPiSetup())
    return wiringPiError(retval)
}

func SetupSys() error {
    retval := int(C.wiringPiSetupSys())
    return wiringPiError(retval)
}

func SetupGpio() int {
    return int(C.wiringPiSetupGpio())
}

func SetupPhys() int {
    return int(C.wiringPiSetupPhys())
}

// WPIPinType
const (
    WPI_PIN_BCM = 1
    WPI_PIN_WPI = 2
    WPI_PIN_PHYS = 3
)

func SetupPinType(pinType int) int {                            // Interface V3.3
    return int(C.wiringPiSetupPinType(C.enum_WPIPinType(pinType)))
}

func SetupGpioDevice(pinType int) int {                         // Interface V3.3
    return int(C.wiringPiSetupGpioDevice(C.enum_WPIPinType(pinType)))
}


/*
extern struct wiringPiNodeStruct *wiringPiFindNode (int pin) ;
extern struct wiringPiNodeStruct *wiringPiNewNode  (int pinBase, int numPins) ;
*/


func GpioDeviceGetFd() int {
    return int(C.wiringPiGpioDeviceGetFd())               // Interface V3.3
}

func PinModeAlt(pin int, mode int) {
    C.pinModeAlt(C.int(pin), C.int(mode))
}

func PinMode(pin int, mode int) {
    C.pinMode(C.int(pin), C.int(mode))
}

func PullUpDnControl(pin int, pud int) {
    C.pullUpDnControl(C.int(pin), C.int(pud))
}

func DigitalRead(pin int) (value int) {
    return int(C.digitalRead(C.int(pin)))
}

func DigitalWrite(pin int, value int) {
    C.digitalWrite(C.int(pin), C.int(value))
}

func PwmWrite(pin int, value int) {
    C.pwmWrite(C.int(pin), C.int(value))
}

func AnalogRead(pin int) (value int) {
    return int(C.analogRead(C.int(pin)))
}

func AnalogWrite(pin int, value int) {
    C.analogWrite(C.int(pin), C.int(value))
}


// Extras from arduino land

func Delay(howLong uint) {
    C.delay(C.uint(howLong))
}

func DelayMicroseconds(howLong uint) {
    C.delayMicroseconds(C.uint(howLong))
}

func Millis() uint {
    return uint(C.millis())
}

func Micros() uint {
    return uint(C.micros())
}

func PiMicros64() uint64 {
    return uint64(C.piMicros64())       // Interface V3.7
}


// On-Board Raspberry Pi hardware specific stuff

func PiGpioLayout() int {
    return int(C.piGpioLayout())
}

func PiBoardRev() int {
    return int(C.piBoardRev())          // Deprecated, but does the same as piGpioLayout
}

func PiBoardId() (model string, revision string, maker string, memsize int, overvolted int) {
    var _model, _rev, _mem, _maker, _overvolted C.int
    C.piBoardId(&_model, &_rev, &_mem, &_maker, &_overvolted)
    model = C.GoString(C.piModelNames[_model])
    revision = C.GoString(C.piRevisionNames[_rev])
    maker = C.GoString(C.piMakerNames[_maker])
    memsize = int(C.piMemorySize[_mem])
    overvolted = int(_overvolted)
    return
}

func PiBoard40Pin() int {
    return int(C.piBoard40Pin())        // Interface V3.7
}

func WpiPinToGpio(wpiPin int) int {
    return int(C.wpiPinToGpio(C.int(wpiPin)))
}

func PhysPinToGpio(physPin int) int {
    return int(C.physPinToGpio(C.int(physPin)))
}

func SetPadDrive(group int, value int) {
    C.setPadDrive(C.int(group), C.int(value))
}

func SetPadDrivePin(pin int, value int) {
    C.setPadDrivePin(C.int(pin), C.int(value))     // Interface V3.0
}

func GetAlt(pin int) int {
    return int(C.getAlt(C.int(pin)))
}

func PwmToneWrite(pin int, freq int) {
    C.pwmToneWrite(C.int(pin), C.int(freq))
}

func PwmSetMode(mode int) {
    C.pwmSetMode(C.int(mode))
}

func PwmSetRange(value uint) {
    C.pwmSetRange(C.uint(value))
}

func PwmSetClock(divisor int) {
    C.pwmSetClock(C.int(divisor))
}

func GpioClockSet(pin int, freq int) {
    C.gpioClockSet(C.int(pin), C.int(freq))
}

func DigitalReadByte() uint {
    return uint(C.digitalReadByte())
}
func DigitalReadByte2() uint {
    return uint(C.digitalReadByte2())
}

func DigitalWriteByte(value int) {
    C.digitalWriteByte(C.int(value))
}
func DigitalWriteByte2(value int) {
    C.digitalWriteByte2(C.int(value))
}
