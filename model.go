package main

type Config struct {
	geminiApiKey          string
	geminiModel           string
	geminiTemperature     float32
	geminiTopK            int32
	geminiTopP            float32
	geminiMaxOutputTokens int32
}
