package handlers

import (
	"fmt"
	"math/rand"
	"server/helpers"
	
)


func getEnabledActions() []int {
	config := helpers.LoadConfig()
	var actions []int







	if config.ErrorInjection.BitFlip {
		actions = append(actions, 1) 
	}
	if config.ErrorInjection.CharacterSubstitution {
		actions = append(actions, 2) 
	}
	if config.ErrorInjection.CharacterDeletion {
		actions = append(actions, 3) 
	}
	if config.ErrorInjection.CharacterInsertion {
		actions = append(actions, 4) 
	}
	if config.ErrorInjection.CharacterSwapping {
		actions = append(actions, 5) 
	}
	if config.ErrorInjection.MultipleBitFlips {
		actions = append(actions, 6) 
	}
	if config.ErrorInjection.BurstError {
		actions = append(actions, 7) 
	}

	if len(actions) == 0 || config.ErrorInjection.NoError {
		actions = append(actions, 0) 
	}

	return actions
}

func injectError(data string) string {
	enabledActions := getEnabledActions()


	action := enabledActions[rand.Intn(len(enabledActions))]

	
	runes := []rune(data)
	if len(runes) == 0 {
		return data
	}

	switch action {
	case 0: // No Error
		fmt.Println("[SERVER-LOG] Durum: Hata Enjekte Edilmedi")

	case 1: // Case Bit Flip
		fmt.Println("[SERVER-LOG] Hata: Bit Flip Enjekte Edildi")
		idx := rand.Intn(len(runes))
		shift := rand.Intn(8) 
		runes[idx] = runes[idx] ^ (1 << shift)

	case 2: // Case Character Substitution
		fmt.Println("[SERVER-LOG] Hata: Karakter Değiştirme Enjekte Edildi")
		idx := rand.Intn(len(runes))
		
		runes[idx] = rune(rand.Intn(26) + 65) 

	case 3: // Case Character Deletion 
		fmt.Println("[SERVER-LOG] Hata: Karakter Silme Enjekte Edildi")
		if len(runes) > 1 {
			idx := rand.Intn(len(runes))
			runes = append(runes[:idx], runes[idx+1:]...)
		}

	case 4: // Case Random Character Insertion
		fmt.Println("[SERVER-LOG] Hata: Karakter Ekleme Enjekte Edildi")
		idx := rand.Intn(len(runes) + 1)    
		newChar := rune(rand.Intn(26) + 65) 

		
		runes = append(runes[:idx], append([]rune{newChar}, runes[idx:]...)...)

	case 5: // Case Character Swapping
		fmt.Println("[SERVER-LOG] Hata: Karakter Kaydırma Enjekte Edildi")
		if len(runes) >= 2 {
			idx := rand.Intn(len(runes) - 1)
			runes[idx], runes[idx+1] = runes[idx+1], runes[idx]
		}

	case 6: // Case Multiple Bit Flips
		fmt.Println("[SERVER-LOG] Hata: Çoklu Bit Flip Enjekte Edildi")
		count := rand.Intn(3) + 2 
		for i := 0; i < count; i++ {
			idx := rand.Intn(len(runes))
			shift := rand.Intn(8)
			runes[idx] = runes[idx] ^ (1 << shift)
		}

	case 7: // Case Burst Error
		fmt.Println("[SERVER-LOG] Hata: Burst Error Enjekte Edildi")
		if len(runes) > 0 {
			burstLen := rand.Intn(6) + 3 
			if burstLen > len(runes) {
				burstLen = len(runes)
			}

			startIdx := rand.Intn(len(runes) - burstLen + 1)
			for i := 0; i < burstLen; i++ {
				runes[startIdx+i] = runes[startIdx+i] ^ 0xFF
			}
		}
	}

	return string(runes)
}
