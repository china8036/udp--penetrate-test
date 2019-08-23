package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	clients := make(map[int][]byte)
	address := "0.0.0.0:10006"
	addr, err := net.ResolveUDPAddr("udp", address)
	checkError(err)

	conn, err1 := net.ListenUDP("udp", addr)
	checkError(err1)
	data := make([]byte, 1024)
	_, rAddr1, err2 := conn.ReadFromUDP(data)
	checkError(err2)
	fmt.Println(fmt.Sprintf("addr 1 is %s:%d", rAddr1.IP.To4().String(),rAddr1.Port))
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, rAddr1.IP)
	binary.Write(bytesBuffer, binary.LittleEndian, uint16(rAddr1.Port))
	binary.Write(bytesBuffer, binary.LittleEndian, []byte(rAddr1.Zone))
	clients[0] = bytesBuffer.Bytes()


	_, rAddr2, err3 := conn.ReadFromUDP(data)
	checkError(err3)
	fmt.Println(fmt.Sprintf("addr 2 is %s:%d", rAddr2.IP.To4().String(),rAddr2.Port))
	bytesBuffer1 := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer1, binary.LittleEndian, rAddr2.IP)
	binary.Write(bytesBuffer1, binary.LittleEndian, uint16(rAddr2.Port))
	binary.Write(bytesBuffer1, binary.LittleEndian, []byte(rAddr2.Zone))
	clients[1] = bytesBuffer1.Bytes()
	fmt.Println(clients[1])
	_, err4 := conn.WriteToUDP(clients[1], rAddr1)
	checkError(err4)
	fmt.Println(clients[0])
	_, err5 := conn.WriteToUDP(clients[0], rAddr2)
	checkError(err5)
	for {
		time.Sleep(1 * time.Second)
	}
}
