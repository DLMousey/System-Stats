package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"tcp-server/lib"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for { // "while"
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn) // Handle the request in a new goroutine
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading:" , err.Error())
		writeAndClose(conn, "Error reading stream: " + err.Error())
	}

	entity := string(buf)

	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		fmt.Println("Failed to compile regex")
		writeAndClose(conn, "Failed to compile regex")
	}

	var sanitisedEntity = reg.ReplaceAllString(entity, "")
	switch sanitisedEntity {
	case "cpu":
		idle0, total0 := lib.CpuLoad()
		time.Sleep(3 * time.Second)
		idle1, total1 := lib.CpuLoad()

		idleTicks := float64(idle1 - idle0)
		totalTicks := float64(total1 - total0)
		cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks
		strCpuUsage := fmt.Sprintf("%f", cpuUsage)

		fmt.Println(strCpuUsage)
		writeAndClose(conn, strCpuUsage)
	case "memory":
		writeAndClose(conn, "TODO: Read free/avail memory")
	default:
		writeAndClose(conn, "Unrecognised command: " + sanitisedEntity)
	}
}

func writeAndClose(conn net.Conn, message string) {
	_, _ = conn.Write([]byte(message))
	_ = conn.Close()
}