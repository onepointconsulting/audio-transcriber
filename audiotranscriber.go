package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type audioMimeType string

type GeminiModel struct {
	Model  *genai.GenerativeModel
	Client *genai.Client
}

const (
	Mp3  audioMimeType = "audio/mpeg"
	Wav  audioMimeType = "audio/wav"
	Ogg  audioMimeType = "audio/ogg"
	M4a  audioMimeType = "audio/mp4"
	Flac audioMimeType = "audio/flac"
	Aac  audioMimeType = "audio/aac"
	Wma  audioMimeType = "audio/x-ms-wma"
)

func getAudioMimeType(ext string) (audioMimeType, bool) {
	switch strings.ToLower(ext) {
	case ".mp3":
		return Mp3, true
	case ".wav":
		return Wav, true
	case ".ogg":
		return Ogg, true
	case ".m4a":
		return M4a, true
	case ".flac":
		return Flac, true
	case ".aac":
		return Aac, true
	case ".wma":
		return Wma, true
	default:
		return "", false
	}
}

func transcribeAudio(path string, recursive bool) {
	ctx := context.Background()
	config := loadConfig()
	gemini := createGeminiModel(ctx, config)
	defer gemini.Client.Close()

	fileURIs, audioFiles := filterFiles(ctx, gemini.Client, path, recursive)

	for i, fileURI := range fileURIs {
		processTranscription(ctx, gemini.Model, i, fileURI, audioFiles[i])
	}
}

func uploadToGemini(ctx context.Context, client *genai.Client, path, mimeType string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	options := genai.UploadFileOptions{
		DisplayName: path,
		MIMEType:    mimeType,
	}
	fileData, err := client.UploadFile(ctx, "", file, &options)
	if err != nil {
		log.Fatalf("Error uploading file: %v", err)
	}

	log.Printf("Uploaded file %s as: %s", fileData.DisplayName, fileData.URI)
	return fileData.URI
}

func createGeminiModel(ctx context.Context, config *Config) *GeminiModel {
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.geminiApiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	model := client.GenerativeModel(config.geminiModel)
	model.SetTemperature(config.geminiTemperature)
	model.SetTopK(config.geminiTopK)
	model.SetTopP(config.geminiTopP)
	model.SetMaxOutputTokens(config.geminiMaxOutputTokens)
	model.ResponseMIMEType = "text/plain"

	return &GeminiModel{
		Model:  model,
		Client: client,
	}
}

func filterFiles(ctx context.Context, client *genai.Client, path string, recursive bool) ([]string, []string) {
	var fileURIs []string
	var audioFiles []string

	if recursive {
		err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if mimeType, ok := getAudioMimeType(filepath.Ext(d.Name())); ok {
				fullPath := filepath.Join(filepath.Dir(path), d.Name())
				audioFiles = append(audioFiles, fullPath)
				fileURIs = append(fileURIs, uploadToGemini(ctx, client, fullPath, string(mimeType)))
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if mimeType, ok := getAudioMimeType(filepath.Ext(f.Name())); ok {
				fullPath := filepath.Join(path, f.Name())
				audioFiles = append(audioFiles, fullPath)
				fileURIs = append(fileURIs, uploadToGemini(ctx, client, fullPath, string(mimeType)))
			}
		}
	}

	return fileURIs, audioFiles
}

func processTranscription(
	ctx context.Context,
	model *genai.GenerativeModel, i int,
	fileURI, audioFile string,
) {
	session := model.StartChat()
	session.History = []*genai.Content{
		{
			Role: "user",
			Parts: []genai.Part{
				genai.FileData{URI: fileURI},
				genai.Text("This file is a recording of a meditation commentary."),
			},
		},
	}

	prompt := `Please transcribe the audio and provide a summary of the content.
The output should be in markdown format with the following sections:
1. Transcription
2. Summary

Here is an example of what the output should look like:

## Transcription:

Settle yourself comfortably. Close your eyes and breathe in deeply. And as you breathe out, let go any tension. Continue to breathe deeply in your own time.

In your mind, say your name, and then tell yourself, relax. Feel yourself relaxing and keep breathing naturally.

Again, in your mind, say your name and say, relax. And feel yourself relaxing and letting go.

Keep breathing naturally. Once more, in your mind, say your name and say, relax. And let go even more.

Ask yourself, how am I feeling?

Whatever you're feeling, just notice it. No need to change it or judge it. Just become aware of how you're feeling.

Enjoy this quiet moment, paying attention to yourself. You are visiting your inner world.

Now let's leave this meditation and go back to the world outside you, the world around you. When you're ready to move on into your day, open your eyes.

## Summary:

This audio file is a guided meditation. It instructs the listener to relax, breathe deeply, and release tension. The meditation encourages self-awareness by prompting the listener to acknowledge their feelings without judgment. It guides the listener to connect with their inner world and then gently brings them back to their surroundings, ready to start their day.
`

	resp, err := session.SendMessage(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Error processing file %s: %v", audioFile, err)
		return
	}

	var sb strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		sb.WriteString(fmt.Sprintf("%v\n", part))
	}

	transcriptionFileName := getTranscriptionFileName(audioFile)
	if err := os.WriteFile(transcriptionFileName, []byte(sb.String()), 0644); err != nil {
		log.Printf("Error writing transcription for %s: %v", audioFile, err)
	}
}

func getTranscriptionFileName(audioFilePath string) string {
	audioFileParent := filepath.Base(audioFilePath)
	directory := path.Dir(filepath.ToSlash(audioFilePath))
	transcriptionFileName := filepath.Join(directory, audioFileParent+".md")
	return transcriptionFileName
}
