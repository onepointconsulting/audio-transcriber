# Audio Transcriber

A command-line tool that uses Google's Gemini AI to transcribe audio files into markdown format. It supports various audio formats and provides accurate transcriptions with summaries.

The audio folder contains some examples of the resulting translations in md files.



## Features

- Transcribes audio files to markdown format
- Supports multiple audio formats (MP3, WAV, OGG, M4A, FLAC, AAC, WMA)
- Recursive directory processing
- Configurable model parameters
- Automatic file type detection
- Markdown-formatted output with transcription and summary sections
- Tries to calculate and store the costs in a JSON file

## Prerequisites

- Go 1.21 or later
- Google Gemini API key

## Building

```bash
go build .
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/audio-transcriber.git
cd audio-transcriber
```

2. Install dependencies:
```bash
go mod download
```

This will produce an executable that you can use.

3. Create a `.env` file in the project root with your Gemini API key:
```env
GEMINI_API_KEY=your_api_key_here
GEMINI_MODEL=gemini-2.0-flash
GEMINI_TEMPERATURE=0.7
GEMINI_TOP_K=40
GEMINI_TOP_P=0.95
GEMINI_MAX_OUTPUT_TOKENS=2048

# Check prices here: https://ai.google.dev/gemini-api/docs/pricing
# Prices in USD per million tokens
GIMINI_INPUT_PRICE=0.70
# Price in USD per million tokens
GIMINI_OUTPUT_PRICE=0.40
```

## Usage

Basic usage:
```bash
audio-transcriber --dir /path/to/audio/files
```

Process files recursively:
```bash
audio-transcriber --dir /path/to/audio/files --recursive
```

Before you can use the executable make sure that you have all the environment variables setup.
The best way would be to create an .env file with all of the variables listed below.

### Command Line Options

- `--dir, -d`: Directory containing audio files (required)
- `--recursive, -r`: Process directories recursively (optional)

## Output Format

The tool generates markdown files with the following structure:

```markdown
## Transcription:

[Full transcription of the audio content]

## Summary:

[Concise summary of the content]
```

## Configuration

The following environment variables can be configured:

- `GEMINI_API_KEY`: Your Google Gemini API key
- `GEMINI_MODEL`: The Gemini model to use (default: gemini-pro)
- `GEMINI_TEMPERATURE`: Model temperature (0.0 to 1.0)
- `GEMINI_TOP_K`: Top-k sampling parameter
- `GEMINI_TOP_P`: Top-p sampling parameter
- `GEMINI_MAX_OUTPUT_TOKENS`: Maximum tokens in the output

## Supported Audio Formats

- MP3 (audio/mpeg)
- WAV (audio/wav)
- OGG (audio/ogg)
- M4A (audio/mp4)
- FLAC (audio/flac)
- AAC (audio/aac)
- WMA (audio/x-ms-wma)

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 