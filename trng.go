// Package trng contains a random number source using
// an Arduino Due connected via USB.
//
// The sketch provided produces random bits, in the form of
// 32bit words, every microsecond and writes them to
// the serial interface at 230,400 baud.
package trng

import (
	"io"

	"github.com/pkg/term"
)

const baud = 230400

// Open opens a connection to the Arduino USB device.
// This device on linux systems is /dev/ttyACM0.
func Open(dev string) (io.ReadCloser, error) {
	return term.Open(dev, term.Speed(baud))
}
