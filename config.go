package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func initDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file: %w", err)
	}
}

func extractFloat(envVar string) float32 {
	str := os.Getenv(envVar)
	assert(str != "", envVar+" is not set")
	price, err := strconv.ParseFloat(str, 32)
	assert(err == nil, envVar+" must be a valid float")
	return float32(price)
}

func loadConfig() *Config {
	initDotEnv()
	geminiApiKey := os.Getenv("GEMINI_API_KEY")
	assert(geminiApiKey != "", "GEMINI_API_KEY is not set")
	geminiModel := os.Getenv("GEMINI_MODEL")
	assert(geminiModel != "", "GEMINI_MODEL is not set")

	geminiTemperature32 := extractFloat("GEMINI_TEMPERATURE")

	topKStr := os.Getenv("GEMINI_TOP_K")
	assert(topKStr != "", "GEMINI_TOP_K is not set")
	geminiTopK, err := strconv.Atoi(topKStr)
	assert(err == nil, "GEMINI_TOP_K must be a valid integer")
	geminiTopK32 := int32(geminiTopK)

	geminiTopP32 := extractFloat("GEMINI_TOP_P")

	maxTokensStr := os.Getenv("GEMINI_MAX_OUTPUT_TOKENS")
	assert(maxTokensStr != "", "GEMINI_MAX_OUTPUT_TOKENS is not set")
	geminiMaxOutputTokens, err := strconv.Atoi(maxTokensStr)
	assert(err == nil, "GEMINI_MAX_OUTPUT_TOKENS must be a valid integer")
	geminiMaxOutputTokens32 := int32(geminiMaxOutputTokens)

	geminiInputPrice32 := extractFloat("GIMINI_INPUT_PRICE")
	geminiOutputPrice32 := extractFloat("GIMINI_OUTPUT_PRICE")

	return &Config{
		geminiApiKey:          geminiApiKey,
		geminiModel:           geminiModel,
		geminiTemperature:     geminiTemperature32,
		geminiTopK:            geminiTopK32,
		geminiTopP:            geminiTopP32,
		geminiMaxOutputTokens: geminiMaxOutputTokens32,
		geminiInputPrice:      geminiInputPrice32,
		geminiOutputPrice:     geminiOutputPrice32,
	}
}
