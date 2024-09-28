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

// Error messages
const (
    ERR_PI_BOARD_NOT_DETECTED = "Pi Board not detected!"
)

// WiringPi modes
const (
    WPI_MODE_PINS             = 0
    WPI_MODE_GPIO             = 1
    WPI_MODE_GPIO_SYS         = 2  // Deprecated since 3.2
    WPI_MODE_PHYS             = 3
    WPI_MODE_PIFACE           = 4
    WPI_MODE_GPIO_DEVICE_BCM  = 5  // BCM pin numbers like WPI_MODE_GPIO
    WPI_MODE_GPIO_DEVICE_WPI  = 6  // WiringPi pin numbers like WPI_MODE_PINS
    WPI_MODE_GPIO_DEVICE_PHYS = 7  // Physic pin numbers like WPI_MODE_PHYS
    WPI_MODE_UNINITIALISED    = -1
)

// Pin modes
const (
    INPUT            = 0
    OUTPUT           = 1
    PWM_OUTPUT       = 2
    PWM_MS_OUTPUT    = 8
    PWM_BAL_OUTPUT   = 9
    GPIO_CLOCK       = 3
    SOFT_PWM_OUTPUT  = 4
    SOFT_TONE_OUTPUT = 5
    PWM_TONE_OUTPUT  = 6
    PM_OFF           = 7   // To input/release line
)

// Pin levels
const (
    LOW              = 0
    HIGH             = 1
)

// Pull up/down/none
const (
    PUD_OFF          = 0
    PUD_DOWN         = 1
    PUD_UP           = 2
)

// PWM modes
const (
    PWM_MODE_MS      = 0
    PWM_MODE_BAL     = 1
)

// Interrupt levels
const (
    INT_EDGE_SETUP    = 0
    INT_EDGE_FALLING  = 1
    INT_EDGE_RISING   = 2
    INT_EDGE_BOTH     = 3
)

// Pi model types and version numbers.
// Intended for the GPIO program Use at your own risk.
// https://www.raspberrypi.com/documentation/computers/raspberry-pi.html#new-style-revision-codes
const (
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
)

// Pi maker numbers.
const (
    PI_MAKER_SONY     = 0
    PI_MAKER_EGOMAN   = 1
    PI_MAKER_EMBEST   = 2
    PI_MAKER_UNKNOWN  = 3
)

// Raspberry Pi GPIO layout numbers.
const (
    GPIO_LAYOUT_PI1_REV1 = 1   // Pi 1 A/B Revision 1, 1.1, CM
    GPIO_LAYOUT_DEFAULT  = 2
)

// BCM to WiringPi pins aliases.
const (
    BCM_GPIO_0  = 30
    BCM_GPIO_1  = 31
    BCM_GPIO_2  = 8
    BCM_GPIO_3  = 9
    BCM_GPIO_4  = 7
    BCM_GPIO_5  = 21
    BCM_GPIO_6  = 22
    BCM_GPIO_7  = 11
    BCM_GPIO_8  = 10
    BCM_GPIO_9  = 13
    BCM_GPIO_10 = 12
    BCM_GPIO_11 = 14
    BCM_GPIO_12 = 26
    BCM_GPIO_13 = 23
    BCM_GPIO_14 = 15
    BCM_GPIO_15 = 16
    BCM_GPIO_16 = 27
    BCM_GPIO_17 = 0
    BCM_GPIO_18 = 1
    BCM_GPIO_19 = 24
    BCM_GPIO_20 = 28
    BCM_GPIO_21 = 29
    BCM_GPIO_22 = 3
    BCM_GPIO_23 = 4
    BCM_GPIO_24 = 5
    BCM_GPIO_25 = 6
    BCM_GPIO_26 = 25
    BCM_GPIO_27 = 2
)

// Native WiringPi pins aliases.
const (
    GPIO_0      = 0
    GPIO_1      = 1
    GPIO_2      = 2
    GPIO_3      = 3
    GPIO_4      = 4
    GPIO_5      = 5
    GPIO_6      = 6
    GPIO_7      = 7
    GPIO_8      = 8
    GPIO_9      = 9
    GPIO_10     = 10
    GPIO_11     = 11
    GPIO_12     = 12
    GPIO_13     = 13
    GPIO_14     = 14
    GPIO_15     = 15
    GPIO_16     = 16
    GPIO_17     = 17
    GPIO_18     = 18
    GPIO_19     = 19
    GPIO_20     = 20
    GPIO_21     = 21
    GPIO_22     = 22
    GPIO_23     = 23
    GPIO_24     = 24
    GPIO_25     = 25
    GPIO_26     = 26
    GPIO_27     = 27
)

// WiringPi I2C pins.
const (
    SDA0        = 30
    SCL0        = 31
    SDA1        = 8
    SCL1        = 9
)

// WiringPi SPI pins.
const (
    CE0         = 10
    CE1         = 11
    MOSI        = 12
    MISO        = 13
    SCLK        = 14
)

// WiringPi UART pins.
const (
    TxD         = 15
    RxD         = 16
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

// WPIPinAlt flags.
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

var _mutex = new(sync.Mutex)

// Checking return value (error code).
func wiringPiError(retval int) error {
    if retval >= 0 {
        return nil
    }
    _mutex.Lock()
    defer _mutex.Unlock()
    return errors.New(C.GoString(C.strerror(C.int(retval))))
}

//
// Wrapper functoins for WiringPi interactions.
//

//
// Core wiringPi functions.
//

// Return WiringPi version number
func Version() (int, int) {
    var major, minor C.int
    C.wiringPiVersion(&major, &minor)
    return int(major), int(minor)
}

// Return global memory access.
func GlobalMemoryAccess() int {
    return int(C.wiringPiGlobalMemoryAccess())  // Interface V3.3
}

// Return user-level access.
func UserLevelAccess() int {
    return int(C.wiringPiUserLevelAccess())
}

// Setup WiringPi library.
// Must be called once at the start of your program execution.
// Default setup: Initialises the system into wiringPi Pin mode and uses the memory mapped hardware directly.
// Changed now to revert to "gpio" mode if we're running on a Compute Module.
func Setup() error {
    retval := int(C.wiringPiSetup())
    return wiringPiError(retval)
}

// Setup WiringPi library (sys).
// GPIO Sysfs Interface for Userspace is deprecated.
// Switched to new GPIO driver Interface in version 3.3.
func SetupSys() error {
    retval := int(C.wiringPiSetupSys())
    return wiringPiError(retval)
}

// Setup WiringPi library (gpio).
// BCM numbering of the GPIOs and directly accesses the GPIO registers.
// Must be called once at the start of your program execution.
// GPIO setup: Initialises the system into GPIO Pin mode and uses the memory mapped hardware directly.
func SetupGpio() int {
    return int(C.wiringPiSetupGpio())
}

// Setup WiringPi library (phys).
// Must be called once at the start of your program execution.
// Phys setup: Initialises the system into Physical Pin mode and uses the memory mapped hardware directly.
func SetupPhys() int {
    return int(C.wiringPiSetupPhys())
}

// WPIPinType
const (
    WPI_PIN_BCM = 1
    WPI_PIN_WPI = 2
    WPI_PIN_PHYS = 3
)

// Setup WiringPi with 'pinType' option:
// - WPI_PIN_BCM calls the WiringPiSetupGpio()
// - WPI_PIN_WPI calls the WiringPiSetup()
// - WPI_PIN_PHYS calls the WiringPiSetupPhys()
func SetupPinType(pinType int) int {                            // Interface V3.3
    return int(C.wiringPiSetupPinType(C.enum_WPIPinType(pinType)))
}

func SetupGpioDevice(pinType int) int {                         // Interface V3.3
    return int(C.wiringPiSetupGpioDevice(C.enum_WPIPinType(pinType)))
}

func GpioDeviceGetFd() int {
    return int(C.wiringPiGpioDeviceGetFd())               // Interface V3.3
}

// This is an un-documented special to let you set any pin to any mode.
func PinModeAlt(pin int, mode int) {
    C.pinModeAlt(C.int(pin), C.int(mode))
}

// Sets the mode of a pin to be input, output or PWM output.
func PinMode(pin int, mode int) {
    C.pinMode(C.int(pin), C.int(mode))
}

// Control the internal pull-up/down resistors on a GPIO pin.
func PullUpDnControl(pin int, pud int) {
    C.pullUpDnControl(C.int(pin), C.int(pud))
}

// Read the value of a given Pin, returning HIGH or LOW.
func DigitalRead(pin int) (value int) {
    return int(C.digitalRead(C.int(pin)))
}

// Set an output bit.
func DigitalWrite(pin int, value int) {
    C.digitalWrite(C.int(pin), C.int(value))
}

// Set an output PWM value.
func PwmWrite(pin int, value int) {
    C.pwmWrite(C.int(pin), C.int(value))
}

// Read the analog value of a given Pin.
// There is no on-board Pi analog hardware, so this needs to go to a new node.
func AnalogRead(pin int) (value int) {
    return int(C.analogRead(C.int(pin)))
}

// Write the analog value to the given Pin.
// There is no on-board Pi analog hardware, so this needs to go to a new node.
func AnalogWrite(pin int, value int) {
    C.analogWrite(C.int(pin), C.int(value))
}


//
// Extras from arduino land.
//

// Wait for some number of milliseconds.
func Delay(howLong uint) {
    C.delay(C.uint(howLong))
}

// This is somewhat intersting. It seems that on the Pi, a single call
// to nanosleep takes some 80 to 130 microseconds anyway, so while
// obeying the standards (may take longer), it's not always what we
// want!
//
// So what I'll do now is if the delay is less than 100uS we'll do it
// in a hard loop, watching a built-in counter on the ARM chip. This is
// somewhat sub-optimal in that it uses 100% CPU, something not an issue
// in a microcontroller, but under a multi-tasking, multi-user OS, it's
// wastefull, however we've no real choice )-:
//
// Plan B: It seems all might not be well with that plan, so changing it
// to use gettimeofday () and poll on that instead...
func DelayMicroseconds(howLong uint) {
    C.delayMicroseconds(C.uint(howLong))
}

// Return a number of milliseconds as an unsigned int.
// Wraps at 49 days.
func Millis() uint {
    return uint(C.millis())
}

// Return a number of microseconds as an unsigned int.
// Wraps after 71 minutes.
func Micros() uint {
    return uint(C.micros())
}

// Return a number of microseconds as an unsigned int64.
func PiMicros64() uint64 {
    return uint64(C.piMicros64())       // Interface V3.7
}

// Return a number of microseconds as an unsigned int64.
// Same as PiMicros64.
func Micros64() uint64 {
    return uint64(C.piMicros64())       // Interface V3.7
}


//
// On-Board Raspberry Pi hardware specific stuff
//


// Return a number representing the hardware revision of the board.
// This is not strictly the board revision but is used to check the layout of the GPIO connector -
// and there are 2 types that we are really interested in here.
// The very earliest Pi's and the ones that came after that which switched some pins...
//
// Revision 1 really means the early Model A and B's.
// Revision 2 is everything else - it covers the B, B+ and CM.
// (... and the Pi 2 - which is a B+ ++  ...)
// (... and the Pi 0 - which is an A+ ...)
//
// The main difference between the revision 1 and 2 system that I use here is the mapping of the GPIO pins.
// From revision 2, the Pi Foundation changed 3 GPIO pins on the (original)
// 26-way header - BCM_GPIO 22 was dropped and replaced with 27, and 0 + 1 - I2C bus 0 was changed to 2 + 3; I2C bus 1.
//
// Additionally, here we set the piModel2 flag too.
// This is again, nothing to do with the actual model, but the major version numbers - the GPIO base
// hardware address changed at model 2 and above (not the Zero though)
func PiGpioLayout() int {
    return int(C.piGpioLayout())
}

func PiBoardRev() int {
    return int(C.piBoardRev())          // Deprecated, but does the same as PiGpioLayout
}

// Return the real details of the board we have.
//
// This is undocumented and really only intended for the GPIO command. Use at your own risk!
// Seems there are some boards with 0000 in them (mistake in manufacture).
//
//      So the distinction between boards that I can see is:
//
//              0000 - Error
//              0001 - Not used
//
//      Original Pi boards:
//              0002 - Model B,  Rev 1,   256MB, Egoman
//              0003 - Model B,  Rev 1.1, 256MB, Egoman, Fuses/D14 removed.
//
//      Newer Pi's with remapped GPIO:
//              0004 - Model B,  Rev 1.2, 256MB, Sony
//              0005 - Model B,  Rev 1.2, 256MB, Egoman
//              0006 - Model B,  Rev 1.2, 256MB, Egoman
//
//              0007 - Model A,  Rev 1.2, 256MB, Egoman
//              0008 - Model A,  Rev 1.2, 256MB, Sony
//              0009 - Model A,  Rev 1.2, 256MB, Egoman
//
//              000d - Model B,  Rev 1.2, 512MB, Egoman (Red Pi, Blue Pi?)
//              000e - Model B,  Rev 1.2, 512MB, Sony
//              000f - Model B,  Rev 1.2, 512MB, Egoman
//
//              0010 - Model B+, Rev 1.2, 512MB, Sony
//              0013 - Model B+  Rev 1.2, 512MB, Embest
//              0016 - Model B+  Rev 1.2, 512MB, Sony
//              0019 - Model B+  Rev 1.2, 512MB, Egoman
//
//              0011 - Pi CM,    Rev 1.1, 512MB, Sony
//              0014 - Pi CM,    Rev 1.1, 512MB, Embest
//              0017 - Pi CM,    Rev 1.1, 512MB, Sony
//              001a - Pi CM,    Rev 1.1, 512MB, Egoman
//
//              0012 - Model A+  Rev 1.1, 256MB, Sony
//              0015 - Model A+  Rev 1.1, 512MB, Embest
//              0018 - Model A+  Rev 1.1, 256MB, Sony
//              001b - Model A+  Rev 1.1, 256MB, Egoman
//
//      A small thorn is the olde style overvolting - that will add in
//              1000000
//
//      The Pi compute module has an revision of 0011 or 0014 - since we only
//      check the last digit, then it's 1, therefore it'll default to not 2 or
//      3 for a Rev 1, so will appear as a Rev 2. This is fine for the most part, but
//      we'll properly detect the Compute Module later and adjust accordingly.
//
// And then things changed with the introduction of the v2...
//
// For Pi v2 and subsequent models - e.g. the Zero:
//
//   [USER:8] [NEW:1] [MEMSIZE:3] [MANUFACTURER:4] [PROCESSOR:4] [TYPE:8] [REV:4]
//   NEW          23: will be 1 for the new scheme, 0 for the old scheme
//   MEMSIZE      20: 0=256M 1=512M 2=1G
//   MANUFACTURER 16: 0=SONY 1=EGOMAN 2=EMBEST
//   PROCESSOR    12: 0=2835 1=2836
//   TYPE         04: 0=MODELA 1=MODELB 2=MODELA+ 3=MODELB+ 4=Pi2 MODEL B 5=ALPHA 6=CM
//   REV          00: 0=REV0 1=REV1 2=REV2
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

// This function is similar to PiBoardId, but returns numeric values instead of strings.
// See PiBoardId function.
func PiBoardIds() (model int, revision int, maker int, memsize int, overvolted int) {
    var _model, _rev, _mem, _maker, _overvolted C.int
    C.piBoardId(&_model, &_rev, &_mem, &_maker, &_overvolted)
    model = int(_model)
    revision = int(_rev)
    maker = int(_maker)
    memsize = int(C.piMemorySize[_mem])
    overvolted = int(_overvolted)
    return
}

// Is a 40-pin Pi board.
func PiBoard40Pin() (bool, error) {
    retval := int(C.piBoard40Pin())        // Interface V3.7
    if retval < 0 {
        return false, errors.New(ERR_PI_BOARD_NOT_DETECTED)
    }
    if retval > 0 {
        return true, nil
    } else {
        return false, nil
    }
}

// Returns the number of GPIO pins of the Raspberry Pi.
// If the Pi Board is not detected, -1 is returned.
func PiBoardPins() (pins int) {
    retval := int(C.piBoard40Pin())        // Interface V3.7
    if retval < 0 {
        return -1
    }
    if retval > 0 {
        pins = 40
    } else {
        pins = 26
    }
    return
}

// Translate a wiringPi Pin number to native GPIO pin number.
// Provided for external support.
func WpiPinToGpio(wpiPin int) int {
    return int(C.wpiPinToGpio(C.int(wpiPin)))
}

// Translate a physical Pin number to native GPIO pin number.
// Provided for external support.
func PhysPinToGpio(physPin int) int {
    return int(C.physPinToGpio(C.int(physPin)))
}

func SetPadDrive(group int, value int) {
    C.setPadDrive(C.int(group), C.int(value))
}

// Set the PAD driver value.
func SetPadDrivePin(pin int, value int) {
    C.setPadDrivePin(C.int(pin), C.int(value))     // Interface V3.0
}

// Returns the ALT bits for a given port.
// Only really of-use for the gpio readall command.
func GetAlt(pin int) int {
    return int(C.getAlt(C.int(pin)))
}

// Pi Specific.
// Output the given frequency on the Pi's PWM pin.
func PwmToneWrite(pin int, freq int) {
    C.pwmToneWrite(C.int(pin), C.int(freq))
}

// Select the native "balanced" mode, or standard mark:space mode
func PwmSetMode(mode int) {
    C.pwmSetMode(C.int(mode))
}

// Set the PWM range register. We set both range registers to the same value.
// If you want different in your own code, then write your own.
func PwmSetRange(value uint) {
    C.pwmSetRange(C.uint(value))
}

// Set/Change the PWM clock.
// Originally WiringPi code, but changed (for the better!) by Chris Hall, <chris@kchall.plus.com>
// after further study of the manual and testing with a scope
func PwmSetClock(divisor int) {
    C.pwmSetClock(C.int(divisor))
}

// Set the frequency on a GPIO clock pin.
func GpioClockSet(pin int, freq int) {
    C.gpioClockSet(C.int(pin), C.int(freq))
}

func DigitalReadByte() uint {
    return uint(C.digitalReadByte())
}
func DigitalReadByte2() uint {
    return uint(C.digitalReadByte2())
}

// Pi Specific.
// Write an 8-bit byte to the first 8 GPIO pins - try to do it as fast as possible.
// However it still needs 2 operations to set the bits, so any external hardware must not
// rely on seeing a change as there will be a change to set the outputs bits to zero,
// then another change to set the 1's.
// Reading is just bit fiddling.
// These are wiringPi pin numbers 0..7, or BCM_GPIO pin numbers:
//   17, 18, 22, 23, 24, 24, 4 on a Pi v1 rev 0-3
//   17, 18, 27, 23, 24, 24, 4 on a Pi v1 rev 3 onwards or B+, 2, 3, zero
func DigitalWriteByte(value int) {
    C.digitalWriteByte(C.int(value))
}

// Pi Specific.
// Write an 8-bit byte to the second set of 8 GPIO pins.
// This is marginally faster than the first lot as these are consecutive BCM_GPIO pin numbers.
// However they overlap with the original read/write bytes.
func DigitalWriteByte2(value int) {
    C.digitalWriteByte2(C.int(value))
}
