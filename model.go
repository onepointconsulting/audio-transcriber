package main

type Config struct {
	geminiApiKey          string
	geminiModel           string
	geminiTemperature     float32
	geminiTopK            int32
	geminiTopP            float32
	geminiMaxOutputTokens int32
	geminiInputPrice      float32
	geminiOutputPrice     float32
}

type CostMetrics struct {
	TotalPromptTokens     int
	TotalCompletionTokens int
	TotalTokens           int
	TotalCost             float32
}
