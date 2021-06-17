package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	setDeviceConfig("192.168.22.149", true)
}



func setDeviceConfig(addr string, isDebug bool) {
	if isDebug {
		fmt.Printf("Sending 60.13.[GS].11.[GS].1 ...")
	}
	conn, err := net.DialTimeout("tcp", addr+":12000", time.Duration(30)*time.Second)
	if err != nil {
		fmt.Printf("Dial error on target %s: %+v\n", addr, err)
		return
	}
	defer conn.Close()

	msg := []byte{2, 0x36, 0x30, 0x2e, 0x31, 0x33,  0x1d, 0x31, 0x31,  0x1d,  0x31, 3, 0}
	lrc := msg[1]
	for i := 2; i < len(msg)-1; i++ {
		lrc = lrc ^ msg[i]
	}
	msg[len(msg)-1] = lrc
	fmt.Printf("sending... - : %+v\n", msg)
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Printf("Error on set device config to %s: %+v\n", addr, err)
		return
	}

	ackBuff := make([]byte, 1)
	conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	_ , err = conn.Read(ackBuff)
	if err != nil {
		checkTimeout(err)
		fmt.Printf("Error reading response from %s: %+v\n", addr, err)
		os.Exit(5)
		return
	}

	if isDebug {
		fmt.Printf("response ackBuff: %+v\n", ackBuff)
	}

	buff := make([]byte, 2000)
	conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	_, err = conn.Read(buff)
	if err != nil {
		return
	}

	n, err := conn.Write([]byte{6})
	if err != nil {
		fmt.Printf("write error: %+v\n", err)
	}

	if isDebug {
		fmt.Printf("response n: %+v\n", n)
		fmt.Printf("response Buff: %+v\n", buff)
	}


	return
}

func checkTimeout(err error) {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		fmt.Printf("Read timed out\n")
		os.Exit(30)
	}
}


//[2 54 48 46 49 51 46 29 46 49 49 46 29 46 49 3 24]
//[STX]60.13[GS]11[GS]1[ETX][CAN]
//[STX]60.13.[GS].11.[GS].1[ETX][CAN]
//6
//0
//.
//1
//3
//.
//gs
//.
//1
//1
//.
//gs
//.
//1
//etx
//can



//2 54 48 46 50 3 25
//  6   0 .   2 stx em