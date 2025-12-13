package main

import (
	"fmt"
	"net"
)

// Parity
func computeParity(text string) string {
	countOnes := 0
	for _, character := range []byte(text) {
		bin := fmt.Sprintf("%08b", character)
		for _, bit := range bin {
			if bit == '1' {
				countOnes++
			}
		}
	}
	if countOnes%2 == 0 {
		return "0"
	}
	return "1"
}

// 2D Parity
func compute2DParity(text string) string {
	for len(text)%8 != 0 {
		text += "\x00"
	}
	rows := len(text) / 8

	rowParity := make([]int, rows)
	columnParity := make([]int, 8)

	charIndex := 0
	for r := 0; r < rows; r++ {
		for columnIndex := 0; columnIndex < 8; columnIndex++ {
			byteValue := text[charIndex]
			bin := fmt.Sprintf("%08b", byteValue)

			for _, bit := range bin {
				if bit == '1' {
					rowParity[r]++
					columnParity[columnIndex]++
				}
			}
			charIndex++
		}
	}

	rowBits := ""
	colBits := ""

	for _, x := range rowParity {
		if x%2 == 0 {
			rowBits += "0"
		} else {
			rowBits += "1"
		}
	}

	for _, x := range columnParity {
		if x%2 == 0 {
			colBits += "0"
		} else {
			colBits += "1"
		}
	}

	return rowBits + colBits
}

// CRC16-CCITT
func computeCRC16(text string) string {
	crc := uint16(0xFFFF)
	for _, character := range []byte(text) {
		crc ^= uint16(character) << 8 //Bu işlem gelen karakteri CRC’nin üst 8 bitine XOR ile ekler
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

// Hamming code
func countBits(h []byte, pos []int) byte {
	count := 0
	for _, p := range pos {
		if h[p] == '1' {
			count++
		}
	}
	if count%2 == 0 {
		return '0'
	}
	return '1'
}

func computeHammingForByte(b byte) string {
	data := fmt.Sprintf("%08b", b)
	h := make([]byte, 13)

	// Veri bitleri Hamming pozisyonlarına yerleştirilir
	h[3] = data[0]
	h[5] = data[1]
	h[6] = data[2]
	h[7] = data[3]
	h[9] = data[4]
	h[10] = data[5]
	h[11] = data[6]
	h[12] = data[7]

	// Parity bitleri hesaplanır
	h[1] = countBits(h, []int{3, 5, 7, 9, 11})
	h[2] = countBits(h, []int{3, 6, 7, 10, 11})
	h[4] = countBits(h, []int{5, 6, 7, 12})
	h[8] = countBits(h, []int{9, 10, 11, 12})

	out := ""
	for i := 1; i <= 12; i++ {
		out += string(h[i])
	}
	return out
}

func computeHamming(text string) string {
	result := ""
	for _, ch := range []byte(text) {
		result += computeHammingForByte(ch) + " "
	}
	return result
}

// Internet checksum
func computeChecksum(text string) string {
	data := []byte(text)
	sum := 0

	if len(data)%2 != 0 {
		data = append(data, 0) // padding
	}

	for i := 0; i < len(data); i += 2 {
		word := int(data[i])<<8 | int(data[i+1])
		sum += word
		if sum > 0xFFFF {
			sum = (sum & 0xFFFF) + 1
		}
	}

	checksum := ^sum & 0xFFFF
	return fmt.Sprintf("%04X", checksum)
}

func main() {
	for {
		var text string
		fmt.Println("[SENDER] Enter a text:")
		fmt.Scanln(&text)

		var choice int
		fmt.Println("[SENDER] Select method: ")
		fmt.Println("1 - Parity")
		fmt.Println("2 - 2D Parity")
		fmt.Println("3 - CRC16")
		fmt.Println("4 - Hamming Code")
		fmt.Println("5 - Internet Checksum")
		fmt.Print("Choice: ")
		fmt.Scanln(&choice)

		methodName := ""
		control := ""

		switch choice {
		case 1:
			methodName = "PARITY"
			control = computeParity(text)

		case 2:
			methodName = "2DPARITY"
			control = compute2DParity(text)

		case 3:
			methodName = "CRC16"
			control = computeCRC16(text)

		case 4:
			methodName = "HAMMING"
			control = computeHamming(text)

		case 5:
			methodName = "CHECKSUM"
			control = computeChecksum(text)

		default:
			fmt.Println("Invalid choice.")
			continue
		}

		packet := fmt.Sprintf("%s|%s|%s\n", text, methodName, control)

		connection, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			fmt.Println("[SENDER] Server'a baglanti basarisiz oldu", err)
			continue
		}
		_, err = connection.Write([]byte(packet))
		if err != nil {
			fmt.Println("[SENDER] Sunucuya mesaj gonderilemedi")
			connection.Close()
			continue
		}

		fmt.Println("\n[SENDER] Gönderilen paket:", packet)
		connection.Close()
	}
}
