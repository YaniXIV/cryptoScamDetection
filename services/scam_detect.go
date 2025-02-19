package services

import (
	"context"
	"cryptoScamDetection/config"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
)

func GenResponse(bodyData string) {
	ctx := context.Background()
	apiKey := config.GetEnv("API_KEY", "noApiKeyProvided")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(bodyData))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
