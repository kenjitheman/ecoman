package openai

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kenjitheman/ecoman/vars"
	openai "github.com/sashabaranov/go-openai"
)

func GenerateAdvice(result string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file (openai.go): %v", err)
		log.Panic(err)
	}
	client := openai.NewClient(os.Getenv("OPENAI_API_TOKEN"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: result + vars.Prompt,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	return strings.TrimSpace(resp.Choices[0].Message.Content)
}
