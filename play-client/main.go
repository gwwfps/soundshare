package main

// #cgo pkg-config: libpulse-simple
/*
 #include <pulse/sample.h>
 #include <pulse/simple.h>
*/
import "C"

import (
	"fmt"
	"net"
	"unsafe"
)

func main() {
	spec := &C.pa_sample_spec{
		format:   C.PA_SAMPLE_FLOAT32,
		channels: 2,
		rate:     48000,
	}

	cerr := C.int(0)

	pa := C.pa_simple_new(
		nil,
		C.CString("soundshare"),
		C.PA_STREAM_PLAYBACK,
		nil,
		C.CString("playback"),
		spec,
		nil,
		nil,
		&cerr,
	)

	fmt.Printf("err: %d\n", cerr)

	addr := net.UDPAddr{
		Port: 30242,
		IP:   net.ParseIP("192.168.1.207"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		p := make([]byte, 5000000)
		s, _, err := ser.ReadFromUDP(p)
		if err != nil {
			fmt.Println(err)
			continue
		}
		go C.pa_simple_write(pa, unsafe.Pointer(&p[0]), C.ulong(s), &cerr)
	}
}
