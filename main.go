package main
import (
	"PortForwarder/libs"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// get the destination url argument from user
	s:=os.Args[1]
	arr := strings.Split(s,".")
	// test if the url is correct lo
	if len(arr)<2 {
		log.Fatalln("url not correct")
	}
	// Listen on all local ports
	listener, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go libs.Handle(conn,s)
	}

}


























// "github.com/akamensky/argparse"

// parser := argparse.NewParser("url", "destination url for the port forwarder")
	// s := parser.String("u", "url", &argparse.Options{Required: true, Help: "true"})
	// err := parser.Parse(os.Args)
	// if err != nil {
	// 	println(parser.Usage(err))
	// }