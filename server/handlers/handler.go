package handlers

import (
	"net"
	"bufio"
	"strings"
	"fmt"
)

func HandleConnection(c1 net.Conn, c2 net.Conn) {
	reader := bufio.NewReader(c1)

	for {
		packet, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		packet = strings.TrimSpace(packet)
		parts := strings.Split(packet, "|")
		if len(parts) < 3 {
			continue
		}

		data := parts[0]
		method := parts[1]
		ctrl := parts[2]

		corruptedData := injectError(data)

		newPacket := fmt.Sprintf("%s|%s|%s\n", corruptedData, method, ctrl)

		_, err = c2.Write([]byte(newPacket))
		if err != nil {
			return
		}

		fmt.Printf("Alıcı Client'a İletilen: %s", newPacket)
	}
}
