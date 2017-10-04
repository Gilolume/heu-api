package main

//Structure Configuration
type Configuration struct {
	GoogleSpeechKey string `json:"google_speech_key"`
}

//--- Struct Response ---
type IndexResponse struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

//--- Fin Struct Response
