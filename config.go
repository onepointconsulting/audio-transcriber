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

func loadConfig() *Config {
	initDotEnv()
	geminiApiKey := os.Getenv("GEMINI_API_KEY")
	assert(geminiApiKey != "", "GEMINI_API_KEY is not set")
	geminiModel := os.Getenv("GEMINI_MODEL")
	assert(geminiModel != "", "GEMINI_MODEL is not set")

	tempStr := os.Getenv("GEMINI_TEMPERATURE")
	assert(tempStr != "", "GEMINI_TEMPERATURE is not set")
	geminiTemperature, err := strconv.ParseFloat(tempStr, 32)
	assert(err == nil, "GEMINI_TEMPERATURE must be a valid float")
	geminiTemperature32 := float32(geminiTemperature)

	topKStr := os.Getenv("GEMINI_TOP_K")
	assert(topKStr != "", "GEMINI_TOP_K is not set")
	geminiTopK, err := strconv.Atoi(topKStr)
	assert(err == nil, "GEMINI_TOP_K must be a valid integer")
	geminiTopK32 := int32(geminiTopK)

	topPStr := os.Getenv("GEMINI_TOP_P")
	assert(topPStr != "", "GEMINI_TOP_P is not set")
	geminiTopP, err := strconv.ParseFloat(topPStr, 32)
	assert(err == nil, "GEMINI_TOP_P must be a valid float")
	geminiTopP32 := float32(geminiTopP)

	maxTokensStr := os.Getenv("GEMINI_MAX_OUTPUT_TOKENS")
	assert(maxTokensStr != "", "GEMINI_MAX_OUTPUT_TOKENS is not set")
	geminiMaxOutputTokens, err := strconv.Atoi(maxTokensStr)
	assert(err == nil, "GEMINI_MAX_OUTPUT_TOKENS must be a valid integer")
	geminiMaxOutputTokens32 := int32(geminiMaxOutputTokens)

	return &Config{
		geminiApiKey:          geminiApiKey,
		geminiModel:           geminiModel,
		geminiTemperature:     geminiTemperature32,
		geminiTopK:            geminiTopK32,
		geminiTopP:            geminiTopP32,
		geminiMaxOutputTokens: geminiMaxOutputTokens32,
	}
}
