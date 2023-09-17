package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func GetChatCompletion(messages []openai.ChatCompletionMessage) *openai.ChatCompletionResponse {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	api_key := os.Getenv("OPENAI_API_KEY")
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	client := openai.NewClient(api_key)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    openai.GPT4,
		Messages: messages,
	})

	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	return &resp
}
