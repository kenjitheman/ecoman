package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func GenerateAdvice(data string) {
	c := openai.NewClient("your token")
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 5,
		Prompt:    data,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("[ERROR] completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
