package trng_test

func ExampleTrngOpen() {
	// open random source
	rand, _ := trng.Open("/dev/ttyACM0")

	// make a tls Config using that source
	config := tls.Config{
		Rand: rand,
	}

	// dial some TLS site
	conn, _ := net.DialTCP("tcp", "www.google.com:443")

	// use our TLS config, with our random source
	// to do the handshake
	client, _ := tls.Client(conn, *config)

	client.Close()
}
