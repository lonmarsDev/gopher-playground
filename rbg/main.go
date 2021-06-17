package main

import (
	"banyanhills.com/revel"
	"banyanhills.com/revel/util"
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func checkTimeout(err error) {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		fmt.Printf("Read timed out\n")

		//	os.Exit(30)
	}
}
func main() {
	localFilename:="E:\\banyanhills\\revel\\DNU_Potts_Test_1_install\\files\\F65235102.OGZ"
	remoteFilename:="F65235102.OGZ"
	target:="192.168.254.102"

	sendFile(localFilename, remoteFilename, target, true, true, 30)
}

func sendFile(localFilename string, remoteFilename string, target string, noreboot bool, debug bool, timeout int) {
	//conn, err := net.Dial("tcp", target+":12000")
	//if err != nil {
	//	fmt.Printf("Dial error on target %s: %+v\n", target, err)
	//	os.Exit(2)
	//	return
	//}
	//defer conn.Close()

	isFirst := true

	file, err := os.Open(localFilename)
	if err != nil {
		fmt.Printf("Error opening file %s: %+v\n", localFilename, err)
		os.Exit(3)
	}
	defer file.Close()

	byteReader := bufio.NewReader(file)
	buff := new(bytes.Buffer)
	//respBuff := []byte{0}

	segmentNumber := 1

	for {
		b, err := byteReader.ReadByte()
		if err != nil {
			break
		}
		if b < 0x20 { // if b less than 32 b =31
			//buff.WriteByte(byte(0xff)  // 255
			buff.WriteByte(byte(b + 0x20))  //32
		} else if b == 0xff {
			buff.WriteByte(byte(0xff))
			buff.WriteByte(byte(0xff))
		} else {
			buff.WriteByte(b)
		}
		if buff.Len() > 3800 {
			if segmentNumber > 99 {
				segmentNumber = 1
			}
			seg := fmt.Sprintf("%02d", segmentNumber)
			segmentNumber += 1
			if isFirst {
				msg := []byte{2, 0x36, 0x32, 0x2e, 0x31, 0x38, seg[0], seg[1], 0x30, 0x30, 0x30, 0x30,
					0x31, 0x31}
				msg = append(msg, []byte(remoteFilename)...)
				msg = append(msg, byte(0x1c))
				msg = append(msg, buff.Bytes()...)
				msg = append(msg, byte(3))
				lrc := msg[1]
				for i := 2; i < len(msg); i++ {
					lrc = lrc ^ msg[i]
				}
				msg = append(msg, byte(lrc))
				//start := 0
				fmt.Printf("Sending segment number %d\n", segmentNumber-1)

				if (segmentNumber-1) == 11 {
					//fmt.Println("conn.Write(msg[start:] : %b" , msg[0])

				}
				//for start < len(msg) {
				//	if debug {
				//		fmt.Printf("Sending segment number %d\n", segmentNumber-1)
				//	}
				//	fmt.Println("con write data ", msg[start:])
				//	//n, err := conn.Write(msg[start:])
				//	//if err != nil {
				//	//	fmt.Printf("Error sending data to %s: %+v\n", target, err)
				//	//	os.Exit(4)
				//	//}
				//	start ++
				//}
				buff = new(bytes.Buffer)
				if debug {
					fmt.Printf("Reading ACK\n")
				}
				fmt.Println("con deadline")
				//conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
				//n, err := conn.Read(respBuff)
				//if err != nil {
				//	checkTimeout(err)
				//	fmt.Printf("Error reading response from %s: %+v\n", target, err)
				//	os.Exit(5)
				//}
				//if n != 1 || respBuff[0] != 6 {
				//	fmt.Printf("No ACK received from %s\n", target, err)
				//	os.Exit(6)
				//}
				isFirst = false
			} else {
				msg := []byte{2, 0x36, 0x32, 0x2e, 0x32, 0x38, seg[0], seg[1], 0x30, 0x30, 0x30, 0x30,
					0x31, 0x31}
				msg = append(msg, buff.Bytes()...)
				msg = append(msg, byte(3))
				lrc := msg[1]
				for i := 2; i < len(msg); i++ {
					lrc = lrc ^ msg[i]
				}
				msg = append(msg, byte(lrc))
				//start := 0
				fmt.Printf("Sending segment number %d\n", segmentNumber-1)
				//will catch 11
				fmt.Println("msg len %v", len(msg) )
				if (segmentNumber-1) == 11 {
					//fmt.Println("conn.Write(msg[start:] : %b" , msg[0:])
					for _, v := range msg {
					fmt.Println( v)
					}
					os.Exit(10)
				}

				//for start < len(msg) {
				//	if debug {
				//		fmt.Printf("Sending segment number %d\n", segmentNumber-1)
				//	}
				//	fmt.Println("conn.Write(msg[start:] :" , msg[start:])
				//	//n, err := conn.Write(msg[start:])
				//	//if err != nil {
				//	//	fmt.Printf("Error sending data to %s: %+v\n", target, err)
				//	//	os.Exit(4)
				//	//}
				//	start ++
				//}
				buff = new(bytes.Buffer)
				if debug {
					fmt.Printf("Reading ACK\n")
				}
				fmt.Println("con set deadline")
				//conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
				//n, err := conn.Read(respBuff)
				//if err != nil {
				//	checkTimeout(err)
				//	fmt.Printf("Error reading response from %s: %+v\n", target, err)
				//	os.Exit(5)
				//}
				//if n != 1 || respBuff[0] != 6 {
				//	fmt.Printf("No ACK received from %s\n", target, err)
				//	os.Exit(6)
				//}
			}
		}
	}
	if segmentNumber > 99 {
		segmentNumber = 1
	}
	seg := fmt.Sprintf("%02d", segmentNumber)
	segmentNumber += 1
	xferType := byte(0x33)
	if isFirst {
		xferType = byte(0x31)
	}
	msg := []byte{2, 0x36, 0x32, 0x2e, xferType, 0x38, seg[0], seg[1], 0x30, 0x30, 0x30, 0x30,
		0x31, 0x31}
	msg = append(msg, buff.Bytes()...)
	msg = append(msg, byte(3))
	lrc := msg[1]
	for i := 2; i < len(msg); i++ {
		lrc = lrc ^ msg[i]
	}
	msg = append(msg, byte(lrc))
	//start := 0
	fmt.Printf("Sending segment number %d\n", segmentNumber-1)
	if (segmentNumber-1) == 11 {
		fmt.Println("conn.Write(msg[start:] : %b" , msg[0])
	}

	//for start < len(msg) {
	//	if debug {
	//		fmt.Printf("Sending segment number %d\n", segmentNumber-1)
	//	}
	//	fmt.Println("conn.Write(msg[start:] :" , msg[start:])
	//	//n, err := conn.Write(msg[start:])
	//	//if err != nil {
	//	//	fmt.Printf("Error sending data to %s: %+v\n", target, err)
	//	//	os.Exit(4)
	//	//}
	//	//start += n
	//}
	if debug {
		fmt.Printf("Reading ACK\n")
	}
	buff = new(bytes.Buffer)
	//conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	//n, err := conn.Read(respBuff)
	//if err != nil {
	//	checkTimeout(err)
	//	fmt.Printf("Error reading response from %s: %+v\n", target, err)
	//	os.Exit(5)
	//}
	//if n != 1 || respBuff[0] != 6 {
	//	fmt.Printf("No ACK received from %s\n", target)
	//	os.Exit(6)
	//}

	resp := make([]byte, 4000)

	if debug {
		fmt.Printf("Reading response\n")
	}
	//conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	//n, err = conn.Read(resp)
	//if err != nil {
	//	checkTimeout(err)
	//	fmt.Printf("Error reading response from %s: %+v\n", target, err)
	//	os.Exit(5)
	//}
	//
	//if n < 7 {
	//	fmt.Printf("Response too short from %s\n", target)
	//	os.Exit(7)
	//}

	switch resp[4] {
	case 0x30:
		fmt.Printf("Transfer successful to %s\n", target)
	case 0x31:
		fmt.Printf("Request out of order error from %s\n", target)
		os.Exit(8)
	case 0x32:
		fmt.Printf("File I/O Error from %s\n", target)
		os.Exit(8)
	case 0x33:
		fmt.Printf("Data Error from %s\n", target)
		os.Exit(8)
	case 0x38:
		fmt.Printf("Error unzipping file error from %s\n", target)
		os.Exit(8)
	case 0x39:
		fmt.Printf("Abort current file error from %s\n", target)
		os.Exit(8)
	default:
		fmt.Printf("Unknown response code %c from %s\n", resp[4], target)
		os.Exit(8)
	}

	if debug {
		fmt.Printf("Sending ACK\n")
	}
	//_, err = conn.Write([]byte{6})
	//if err != nil {
	//	fmt.Printf("Error sending ACK to %s: %+v\n", target, err)
	//	os.Exit(9)
	//}

	ucfilename := strings.ToUpper(remoteFilename)

	if !noreboot && (strings.HasSuffix(ucfilename, ".OGZ") ||
		strings.HasSuffix(ucfilename, ".TGZ") ||
		strings.HasSuffix(ucfilename, ".K3Z")) {

		msg = []byte{2, 0x39, 0x37, 0x2e, 3, 0}
		lrc = msg[1]

		fmt.Println("lcr :%s", lrc)
		for i := 2; i < len(msg)-1; i++ {
			lrc = lrc ^ msg[i]
		}
		msg[len(msg)-1] = lrc
		fmt.Printf("Sending reboot request\n")
		//_, err = conn.Write(msg)
		//if err != nil {
		//	fmt.Printf("Error sending reboot request to %s: %+v\n", target, err)
		//	os.Exit(11)
		//}
		fmt.Printf("Reboot Request Sent to %s\n", target)
	}
}
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
func check_subnet(excludeIPS []string, addrChan, addrDoneChan, resultChan chan string, handlerFunc func(string) string) {

	excludeMap := map[string]bool{}
	for _, e := range excludeIPS {
		excludeMap[e] = true
	}
	for {
		addr := <-addrChan
		if addr == "" {
			return
		}
		_, ok := excludeMap[addr]
		if !ok {
			result := handlerFunc(addr)
			if result != "" {
				resultChan <- result
			}
		}

		addrDoneChan <- ""
	}
}

func forallPaymentTerminals(subnets []string, excludeIPS []string, handlerFunc func(string) string) []string {
	results := []string{}
	for _, subnet := range subnets {
		hosts, err := Hosts(subnet)
		if err != nil {
			fmt.Printf("Error parsing subnet: %s\n", subnet)
			os.Exit(13)
		}

		addrChan := make(chan string, 254+20)
		addrDoneChan := make(chan string)
		resultChan := make(chan string)
		resultsChan := make(chan bool)
		numAddrs := 0

		go (func() {
			for {
				result := <-resultChan
				if result == "" {
					resultsChan <- true
					break
				}
				results = append(results, result)
			}
		})()

		for i := 0; i < 20; i++ {
			go check_subnet(excludeIPS, addrChan, addrDoneChan, resultChan, handlerFunc)
		}
		for _, ip := range hosts {
			addrChan <- ip
			numAddrs += 1
		}
		for i := 0; i < 20; i++ {
			addrChan <- ""
		}

		for i := 1; i <= numAddrs; i++ {
			_ = <-addrDoneChan
		}

		go (func() { resultChan <- "" })()
		<-resultsChan
	}
	return results
}

func splitRBAReply(buff []byte) [][]byte {
	splits := [][]byte{}
	start := 0
	end := 1
	for start < len(buff) && end <= len(buff) {
		if end == len(buff) {
			splits = append(splits, buff[start:])
			break
		} else if buff[end] == 0x1c {
			splits = append(splits, buff[start:end])
			start = end + 1
			end = start + 1
		} else {
			end += 1
		}
	}
	return splits
}

func getSerialNumber(addr string, serialSearch map[string]bool, debug bool) string {
	conn, err := net.DialTimeout("tcp", addr+":12000", time.Duration(30)*time.Second)
	if err != nil {
		return ""
	}
	defer conn.Close()

	msg := []byte{2, 0x30, 0x38, 0x2e, 0x30, 3, 0}
	lrc := msg[1]
	for i := 2; i < len(msg)-1; i++ {
		lrc = lrc ^ msg[i]
	}
	msg[len(msg)-1] = lrc
	_, err = conn.Write(msg)
	if err != nil {
		return ""
	}

	ackBuff := make([]byte, 1)
	conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	n, err := conn.Read(ackBuff)
	if err != nil {
		return ""
	}

	if ackBuff[0] != 6 {
		return ""
	}

	buff := make([]byte, 2000)
	conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	n, err = conn.Read(buff)
	if err != nil {
		return ""
	}

	_, err = conn.Write([]byte{6})
	if err != nil {
		return ""
	}

	if n < 4 {
		return ""
	}
	if buff[1] != 0x30 || buff[2] != 0x38 || buff[3] != 0x2e {
		return ""
	}

	splits := splitRBAReply(buff[4 : n-2])

	if len(splits) < 23 {
		return ""
	}

	ser := string(splits[7])
	_, ok := serialSearch[ser]
	if ok {
		return addr
	} else {
		return ""
	}
}

func findSerials(scanSerial []string, manifest revel.RevelManifest, debug, scanStatic bool) []string {
	excludeIPS := []string{}
	if !scanStatic {
		for _, device := range manifest.Devices {
			if device.IP != "" {
				excludeIPS = append(excludeIPS, device.IP)
			}
		}
	}

	serialMap := map[string]bool{}
	for _, ser := range scanSerial {
		serialMap[ser] = true
	}
	return forallPaymentTerminals(getScanAddresses(debug), excludeIPS, func(addr string) string {
		return getSerialNumber(addr, serialMap, debug)
	})
}

func getScanAddresses(debug bool) []string {
	scanAddresses := []string{"192.168.22.0/24"}
	cfg, err := revel.LoadConfig()
	if err != nil {
		if debug {
			fmt.Printf("Error loading leaf_services.json: %+v\n", err)
		}
		return scanAddresses
	}

	addrList := []string{}
	for _, s := range cfg.PaymentTerminalMonitor.Subnets {
		addrs := util.ExpandSubnet(s, debug)
		addrList = append(addrList, addrs...)
	}
	return addrList
}
