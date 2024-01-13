package main

import (
	"fmt"
	"os"

	"github.com/akagami-harsh/GoOpenAI/internal/generator"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	organization := os.Getenv("API_ORG")

	client := generator.NewClient(apiKey, organization)

	r := generator.CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []generator.Message{
			{
				Role:    "user",
				Content: "Say this is a test!",
			},
		},
		Temperature: 0.7,
	}

	completions, err := client.CreateCompletions(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(completions)
	/* Response should be like this
	{
	  "id": "chatcmpl-xxx",
	  "object": "chat.completion",
	  "created": 1678667132,
	  "model": "gpt-3.5-turbo-0301",
	  "usage": {
	    "prompt_tokens": 13,
	    "completion_tokens": 7,
	    "total_tokens": 20
	  },
	  "choices": [
	    {
	      "message": {
	        "role": "assistant",
	        "content": "\n\nThis is a test!"
	      },
	      "finish_reason": "stop",
	      "index": 0
	    }
	  ]
	}
	*/
}
