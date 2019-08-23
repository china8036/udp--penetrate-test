package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

const SERVER_RECV_LEN = 1024

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "120.xx.xxx.xxx:10006")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 10008} // 注意端口必须固定
	udpCon, err := net.DialUDP("udp", srcAddr, serverAddr)
	checkError(err)
	toWrite := "hello"
	defer udpCon.Close()
	fmt.Println("Write:", toWrite)
	_, err = udpCon.Write([]byte(toWrite))
	checkError(err)
	msg := make([]byte, SERVER_RECV_LEN)
	_, err = udpCon.Read(msg)
	fmt.Println(msg)
	checkError(err)
	ip := msg[0:16]
	port := binary.LittleEndian.Uint16(msg[16:18])
	fmt.Println(ip, port)

	udpAddr := net.UDPAddr{ip, int(port), ""}
	udpCon.Close()
	udpCon, err = net.DialUDP("udp", srcAddr, &udpAddr)
	checkError(err)
	go func() {
		for {
			msg := make([]byte, SERVER_RECV_LEN)
			_, err = udpCon.Read(msg)
			checkError(err)
			fmt.Println("Response:", string(msg))
		}
	}()

	for {
		fmt.Println("Send:", toWrite)
		_, err3 := udpCon.Write([]byte(toWrite))
		checkError(err3)
		time.Sleep(1 * time.Second)
	}
}
