package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("E:\\banyanhills\\revel\\DNU_Potts_Test_1_install\\files\\F65235102.OGZ")
	if err != nil {
		fmt.Printf("Error opening file %s: %+v\n", "theLocalFile", err)
		os.Exit(3)
	}
	defer file.Close()

	byteReader := bufio.NewReader(file)
	buff := new(bytes.Buffer)
	//respBuff := []byte{0}

	b, err := byteReader.ReadByte()
	if err != nil {
		panic(err)
	//	break
	}
	fmt.Println(0x20) // 32
	fmt.Println(0xff) //255
	fmt.Println("the b ",b)

	buff.WriteByte(byte(b + 0x20))  //32
	segmentNumber := 1
	seg := fmt.Sprintf("%02d", segmentNumber)
	fmt.Println(seg)
	segmentNumber += 1

	msg := []byte{2, 0x36, 0x32, 0x2e, 0x32, 0x38, seg[0], seg[1], 0x30, 0x30, 0x30, 0x30,
		0x31, 0x31}
	fmt.Println(0x38)
	fmt.Println(string(msg[1:]))



	msg = append(msg, []byte("F65235102.OGZ")...)
	msg 	= append(msg, byte(0x1c))
	msg = append(msg, buff.Bytes()...)
	msg = append(msg, byte(3))
	lrc := msg[1]
	for i := 2; i < len(msg); i++ {
		lrc = lrc ^ msg[i]
	}
	msg = append(msg, byte(lrc))
	fmt.Println( msg)

	fmt.Println("reboot message")
	msg = []byte{2, 0x39, 0x37, 0x2e, 3, 0}
	//fmt.Printf("%s", msg)
	fmt.Println( string(msg))
	fmt.Println( string(0x39))
	fmt.Println( string(0x37))
	fmt.Println( string(0x2e))
	fmt.Println( string(0x2e))
	lrc = msg[1]
	fmt.Println("lcr :%s", lrc)
	for i := 2; i < len(msg)-1; i++ {
		lrc = lrc ^ msg[i]
		fmt.Println( "lrc")
		fmt.Println( string(lrc))
	}
	msg[len(msg)-1] = lrc
	fmt.Println("the loop")
	for _, v := range msg {
		fmt.Println("%s", v)
	}
	fmt.Println("end the loop")
	fmt.Println( string( msg ) )


	fmt.Printf("Sending reboot request\n")

	fmt.Printf("%s", msg)


	fmt.Println("the serial number")
	msg = []byte{2, 0x30, 0x38, 0x2e, 0x30, 3, 0}
	fmt.Println(string(msg[0]))
	fmt.Println(string(msg[1]))
	lrc = msg[1]
	for i := 2; i < len(msg)-1; i++ {
		lrc = lrc ^ msg[i]
	}
	msg[len(msg)-1] = lrc
	fmt.Println(string(msg[0]))


	//var s string

	//...



	fmt.Println("%b", []byte("8"))


	fmt.Println( string("0x1c"))

	//fs 0x1C
	//gs 0x1D

	//60.13.[GS].11.[GS].1  //response is 60.2 means success

	//0x36
     //0x30
     //0x2e  -- .
     //0x31  -- 1
     //0x33
     //0x2e
     //0x1D  --GS
     //0x2e
     //0x31
     //0x31
     //0x2e
	 //0x1D  --GS
     //0x2e  -- .
	 //0x31  -- 1

	//<2020-11-17 09:41:52.124> [TRACE]:{56880}:MessageManager: Send Configuration Write
	//<2020-11-17 09:41:52.124> [TRACE]:{56880}:PacketManager: RBA SDK ==> RBA (13) [STX]60.13[GS]11[GS]1[ETX][CAN]
	//<2020-11-17 09:41:52.149> [TRACE]:{56880}:PacketManager: RBA SDK <== RBA (1) ACK
	//<2020-11-17 09:41:52.262> [TRACE]:{39008}:PacketManager: RBA SDK <== RBA (7) [STX]60.2[ETX][EM]
	//<2020-11-17 09:41:52.262> [TRACE]:{39008}:PacketManager: RBA SDK ==> RBA (1) ACK


	msg = []byte{2, 0x36, 0x30, 0x2e, 0x31, 0x33, 0x2e, 0x1d, 0x2e, 0x31, 0x31, 0x2e, 0x1d, 0x2e, 0x31, 0X3, 0}
	fmt.Println( string(msg))
	lrc = msg[1]
	for i := 2; i < len(msg)-1; i++ {
		lrc = lrc ^ msg[i]
	}
	msg[len(msg)-1] = lrc
	fmt.Println( string(msg))
}


//func


//[2 54 48 46 57 3 18
   2 54 48 46 57 3 18
 //    6  0   . 9  etc dc2