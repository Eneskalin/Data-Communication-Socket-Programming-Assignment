package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)


func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("9000 portu dinlenemedi:", err)
		return
	}
	defer listener.Close()

	fmt.Println("[ClientReceiver] 9000 portunda dinleniyor...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("[ClientReceiver] Server bağlandı.")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		packet, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		packet = strings.TrimSpace(packet)
		fmt.Println("\n[ClientReceiver] Ham Paket:", packet)

		parts := strings.SplitN(packet, "|", 3)
		if len(parts) != 3 {
			fmt.Println("Geçersiz paket formatı")
			continue
		}

		data := parts[0]
		method := parts[1]
		control := parts[2]

		fmt.Println("DATA   :", data)
		fmt.Println("METHOD :", method)
		fmt.Println("CONTROL:", control)

		var calculated string

		switch method {
		case "PARITY":
			calculated = computeParity(data)

		case "2DPARITY":
			calculated = compute2DParity(data)

		case "CRC16":
			calculated = computeCRC16(data)

		case "HAMMING":
			calculated = computeHamming(data)

		case "CHECKSUM":
			calculated = computeChecksum(data)

		default:
			fmt.Println("Bilinmeyen metod")
			continue
		}

		fmt.Println("HESAPLANAN:", calculated)

		if calculated == control {
			fmt.Println("SONUÇ : DATA CORRECT")
		} else {
			fmt.Println("SONUÇ : DATA CORRUPTED")
		}

		fmt.Println("----------------------------")
	}
}


func computeParity(text string) string {
	count := 0
	for _, c := range []byte(text) {
		for i := 0; i < 8; i++ {
			if (c>>i)&1 == 1 {
				count++
			}
		}
	}
	if count%2 == 0 {
		return "0"
	}
	return "1"
}

func compute2DParity(text string) string {
	for len(text)%8 != 0 {
		text += "\x00"
	}
	rows := len(text) / 8
	rowParity := make([]int, rows)
	colParity := make([]int, 8)

	idx := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < 8; c++ {
			b := text[idx]
			for i := 0; i < 8; i++ {
				if (b>>i)&1 == 1 {
					rowParity[r]++
					colParity[i]++
				}
			}
			idx++
		}
	}

	result := ""
	for _, r := range rowParity {
		if r%2 == 0 {
			result += "0"
		} else {
			result += "1"
		}
	}
	for _, c := range colParity {
		if c%2 == 0 {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func computeCRC16(text string) string {
	crc := uint16(0xFFFF)
	for _, ch := range []byte(text) {
		crc ^= uint16(ch) << 8
		for i := 0; i < 8; i++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}
	return fmt.Sprintf("%04X", crc&0xFFFF)
}

func computeHamming(text string) string {
	result := ""
	for _, ch := range []byte(text) {
		result += fmt.Sprintf("%08b ", ch)
	}
	return strings.TrimSpace(result)
}

func computeChecksum(text string) string {
	data := []byte(text)
	sum := 0

	if len(data)%2 != 0 {
		data = append(data, 0)
	}

	for i := 0; i < len(data); i += 2 {
		word := int(data[i])<<8 | int(data[i+1])
		sum += word
		if sum > 0xFFFF {
			sum = (sum & 0xFFFF) + 1
		}
	}

	check := ^sum & 0xFFFF
	return fmt.Sprintf("%04X", check)
}
