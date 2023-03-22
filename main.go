package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/urfave/cli"
)

func getTokenFromEnv() string {
	if os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}
	return os.Getenv("OPENAI_API_KEY")
}

func main() {
	app := cli.NewApp()
	app.Name = "geeptie"
	app.Usage = "GPT Geenie"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "max-tokens, m",
			Usage: "Maximum number of tokens to generate",
		},
		cli.Float64Flag{
			Name:  "temperature, T",
			Usage: "Temperature for the GPT-3 model",
			Value: 0.1,
		},
		cli.Float64Flag{
			Name:  "top-p, p",
			Usage: "Top-p for the GPT-3 model",
		},
		cli.IntFlag{
			Name:  "n",
			Usage: "Number of completions to generate",
			Value: 1,
		},
		cli.BoolFlag{
			Name:  "stream, s",
			Usage: "Stream the output",
		},
		// list models
		cli.BoolFlag{
			Name:  "list-models, L",
			Usage: "List all models",
		},
		cli.GenericFlag{
			Name:  "model, M",
			Usage: "Model to use",
			Value: &EnumValue{
				Enum:    AllModels,
				Default: GPT3Dot5Turbo,
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("list-models") {
			for _, model := range AllModels {
				fmt.Println(model)
			}
			return nil
		}
		maxTokens := c.Int("max-tokens")
		temperature := float32(c.Float64("temperature"))
		topP := float32(c.Int("top-p"))
		n := c.Int("n")
		stream := c.Bool("stream")
		prompt := c.Args()
		fmt.Println("tail", prompt)

		client := openai.NewClient(getTokenFromEnv())
		ctx := context.Background()
		completion := openai.CompletionRequest{
			Prompt:      strings.Join(prompt, " "),
			MaxTokens:   maxTokens,
			Temperature: temperature,
			TopP:        topP,
			N:           n,
			Stream:      stream,
			Model:       openai.GPT3Ada,
		}
		resp, err := client.CreateCompletion(ctx, completion)
		if err != nil {
			log.Fatal(err)
		}
		for i, choice := range resp.Choices {
			fmt.Printf("Choice %d:\n%s\n", i, choice.Text)
		}

		return nil

	}

	app.Run(os.Args)

}
