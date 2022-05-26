// package portForwarder

// import (
// 	"io"
// 	"log"
// 	"net"
// )

// func Handle(src net.Conn) {
// 	dst, err := net.Dial("tcp", "joescatcam.website:80")
// 	if err != nil {
// 	log.Fatalln("Unable to connect to our unreachable host")
// 	}
// 	defer dst.Close()
// 	// Run in goroutine to prevent io.Copy from blocking
// 	 go func() {
// 	// Copy our source's output to the destination
// 	if _, err := io.Copy(dst, src); err != nil {
// 	log.Fatalln(err)
// 	}
// 	}()
// 	// Copy our destination's output back to our source
// 	if _, err := io.Copy(src, dst); err != nil {
// 	log.Fatalln(err)
// 	}
// 	}
package libs

import (
	"io"
	"log"
	"net"
	
	
)

func Handle(src net.Conn,dstsktaddr string) {
	
	dst, err := net.Dial("tcp", dstsktaddr)
	
	if err != nil {
	log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()
	// Run in goroutine to prevent io.Copy from blocking
	 go func() {
	// Copy our source's output to the destination
	if _, err := io.Copy(dst, src); err != nil {
	log.Fatalln(err)
	}
	}()
	// Copy our destination's output back to our source
if _, err := io.Copy(src, dst); err != nil {
	log.Fatalln(err)
	}
	}
	