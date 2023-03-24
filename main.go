package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

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
	app.Name = "qwe"
	app.Usage = "GPT Querying"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "max-tokens, m",
			Usage: "Maximum number of tokens to generate",
			Value: 10,
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
			Name:  "number-completions-n",
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

	app.Action = func(clictx *cli.Context) error {
		if clictx.Bool("list-models") {
			for _, model := range AllModels {
				fmt.Println(model)
			}
			return nil
		}
		maxTokens := clictx.Int("max-tokens")
		temperature := float32(clictx.Float64("temperature"))
		topP := float32(clictx.Int("top-p"))
		n := clictx.Int("n")
		doStream := clictx.Bool("stream")
		prompt := clictx.Args()
		fmt.Println("tail", prompt)

		c := openai.NewClient(getTokenFromEnv())
		ctx := context.Background()
		req := openai.CompletionRequest{
			MaxTokens:   maxTokens,
			Temperature: temperature,
			TopP:        topP,
			N:           n,
			Stream:      doStream,
			Model:       openai.GPT3Ada,
		}

		if doStream {

			stream, err := c.CreateCompletionStream(ctx, req)
			if err != nil {
				return err
			}

			defer stream.Close()

			for {
				resp, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					fmt.Println("Stream finished")
					return nil
				}
				if err != nil {
					return err
				}
				cs := resp.Choices
				if len(cs) > 1 {
					return fmt.Errorf("got more than one choice in stream: %v", cs)
				}
				fmt.Print(cs[0].Text)
			}
		}

		resp, err := c.CreateCompletion(ctx, req)
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
