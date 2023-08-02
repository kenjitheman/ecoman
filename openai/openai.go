package openai

import (
	"context"
	"fmt"
	// "log"
	"os"
	"strings"

	// "github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func GenerateAdvice(result string) string {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	fmt.Printf("[ERROR] error loading .env file: %v", err)
	// 	log.Fatal("[ERROR] error loading .env file")
	// }
	client := openai.NewClient(os.Getenv("OPENAI_APITOKEN"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: result + " here is data, please, can you generate some advices what is best to do on this day, based on data you got",
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
