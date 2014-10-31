trng
----

Adapted from [http://www.kerrywong.com/2014/10/19/using-arduino-dues-true-random-number-generator/](www.kerrywong.com/2014/10/19/using-arduino-dues-true-random-number-generator/)

`trng` is a small Go package that provides an `io.ReadCloser` of random entropy which can be used for cryptographic purposes.

The project is based on the work of Kerry D. Wong, who demonstrated how to extract the random data from the onboard TRNG generator on the SAM3X chip that powers the Arduino Due.

The major change from Wong's work is the raw 32bit output of the TRNG is transmitted to the host and can be directly used as a source of random entropy (Wong's work transmitted the data as string representation of a 32bit signed integer). 

documentation
-------------

[https://godoc.org/github.com/davecheney/trng](godoc.org/github.com/davecheney/trng)

performance
-----------

The serial connection between the atmel and the host computer runs at 230,400 baud, which is slightly faster than the SAM3X can produce random data. In benchmarking entropy is produced at a rate of 18kB/sec.

	% random-bits | dd of=/dev/null bs=1k count=1k iflag=fullblock
	1024+0 records in
	1024+0 records out
	1048576 bytes (1.0 MB) copied, 58.1198 s, 18.0 kB/s

usage
-----

1. Load the `ardunio-due-trng` sketch onto your Ardunio Due. It has to be a SAM3X powered device, older 8bit Atmega Arduinos do not have the hardware trng device.
2. Either use this package, see the [https://godoc.org/github.com/davecheney/trng#example-Open](example in godoc), or use the supplied `random-bits` example program.

   % random-bits | dieharder -a -g 200

security
--------

"Is this really secure? Can I trust you?"

No, this probably isn't secure. No, you shouldn't trust me, it's possible that i've screwed up and the random data isn't actually that random.

