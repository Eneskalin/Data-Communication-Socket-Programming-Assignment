package modal

type ErrorInjectionConfig struct {
	Ports struct {
		SenderPort   string `json:"sender_port"`
		ReceiverPort string `json:"receiver_port"`
	} `json:"ports"`
	ErrorInjection struct {
		BitFlip               bool `json:"bitFlip"`
		CharacterSubstitution bool `json:"characterSubstitution"`
		CharacterInsertion    bool `json:"characterInsertion"`
		CharacterAddition     bool `json:"characterAddition"`
		CharacterDeletion     bool `json:"characterDeletion"`
		CharacterSwapping     bool `json:"characterSwapping"`
		MultipleBitFlips      bool `json:multipleBitFlips`
		BurstError            bool `json:burstError`
		NoError               bool `json:"noError"`
	} `json:"errorInjection"`
}
