package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	AudioDir string
)
var (
	Recursive bool
)

var rootCmd = &cobra.Command{
	Use:   "audio-transcriber",
	Short: "A tool to transcribe audio files using Gemini",
	Long: `A tool that uses Google's Gemini AI to transcribe audio files.
It supports various audio formats and provides accurate transcriptions.`,
	Run: func(cmd *cobra.Command, args []string) {
		if AudioDir == "" {
			fmt.Println("Error: --dir flag is required")
			cmd.Help()
			os.Exit(1)
		}
		// TODO: Add your transcription logic here
		fmt.Printf("Transcribing file: %s\n", AudioDir)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&AudioDir, "dir", "d", "", "Audio directory to transcribe (required)")
	rootCmd.MarkFlagRequired("dir")
	rootCmd.Flags().BoolVarP(&Recursive, "recursive", "r", false, "Transcribe all files in the directory recursively")
}

func Execute() error {
	return rootCmd.Execute()
}
