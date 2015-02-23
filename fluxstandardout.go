// Copyright (c) 2015 Joshua Scoggins
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgement in the product documentation would be
//    appreciated but is not required.
// 2. Altered source versions must be plainly marked as such, and must not be
//    misrepresented as being the original software.
// 3. This notice may not be removed or altered from any source distribution.
//

// sends standard input over the network (similar to netcat but this is how I learn)
package main

import (
	"flag"
	"fmt"
	"github.com/DrItanium/neuron"
	"log"
	"net"
)

var host = flag.String("host", "", "host to connect to")
var port = flag.Uint("port", 2000, "port to connect to")

func main() {
	flag.Parse()
	str := fmt.Sprintf("%s:%d", *host, *port)
	conn, err := net.Dial("tcp", str)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	input := neuron.NewStandardInReader()
	for {
		datum, err := input.ReadByte()
		if err != nil {
			break
		}
		_, err2 := fmt.Fprintf(conn, "%c", datum)
		if err2 != nil {
			log.Print(err)
			break
		}
	}
}
