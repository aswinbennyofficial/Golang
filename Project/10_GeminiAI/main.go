package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	GEMINI_API_KEY := os.Getenv("GEMINI_API_KEY")
	// fmt.Println(GEMINI_API_KEY)

	// Initialise Generative model
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("Write a story starting with - once upon a time in hollywood"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*&resp.Candidates[0].Content.Parts[0])
	//fmt.Printf("%+v", resp)
}
