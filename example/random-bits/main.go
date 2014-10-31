// Provides a stream of random bits suitable for piping to other
// consumers of random data.
package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/davecheney/trng"
)

var dev = flag.String("dev", "/dev/ttyACM0", "arduino device")

func main() {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	flag.Parse()
	rand, err := trng.Open(*dev)
	check(err)
	_, err = io.Copy(os.Stdout, rand)
	check(err)
}
