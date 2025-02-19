package services

import (
	"context"
	"cryptoScamDetection/config"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"strings"
)

func GenResponse(bodyData string) string {
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

	return extractText(resp)
}

func extractText(resp *genai.GenerateContentResponse) string {
	var sb strings.Builder
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if text, ok := part.(genai.Text); ok {
					sb.WriteString(string(text))
				}
			}
		}
	}
	return sb.String()
}
