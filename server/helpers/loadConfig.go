package helpers

import(
	"server/modal"
	"fmt"
	"os"
	"encoding/json"
)



func LoadConfig() modal.ErrorInjectionConfig {
	var config modal.ErrorInjectionConfig
	data, err := os.ReadFile("config/config.json")
	if err != nil {
		fmt.Println("[ERROR] Config dosyası okunamadı:", err)
		return config
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("[ERROR] Config dosyası parse edilemedi:", err)
		return config
	}

	return config
}