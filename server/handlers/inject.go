package handlers

import (
	"fmt"
	"math/rand"
	"time"
)

func injectError(data string) string {
	rand.Seed(time.Now().UnixNano())
	action := rand.Intn(4) // 0-3 arası rastgele sayı

	runes := []rune(data)
	if len(runes) == 0 {
		return data
	}

	switch action {
	case 0: // Bit Flip (Basitçe bir karakteri değiştiriyoruz simülasyon için)
		fmt.Println("[LOG] Hata: Karakter Değiştirme")
		idx := rand.Intn(len(runes))
		runes[idx] = 'X'
	case 1: // Karakter Ekleme
		fmt.Println("[LOG] Hata: Karakter Ekleme")
		idx := rand.Intn(len(runes))
		// Slice arasına karakter sokma işlemi
		runes = append(runes[:idx+1], runes[idx:]...)
		runes[idx] = '?'
	case 2: // Olduğu gibi gönder (Hata yok)
		fmt.Println("[LOG] Hata Yok")
	case 3: // Karakter Silme
		fmt.Println("[LOG] Hata: Karakter Silme")
		if len(runes) > 1 {
			idx := rand.Intn(len(runes))
			runes = append(runes[:idx], runes[idx+1:]...)
		}
	}
	return string(runes)
}
