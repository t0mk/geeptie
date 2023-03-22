package main

import (
	"fmt"
	"strings"
)

type EnumValue struct {
	Enum     []string
	Default  string
	selected string
}

func (e *EnumValue) Set(value string) error {
	for _, enum := range e.Enum {
		if enum == value {
			e.selected = value
			return nil
		}
	}

	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", "))
}

func (e EnumValue) String() string {
	if e.selected == "" {
		return e.Default
	}
	return e.selected
}

const (
	GPT432K0314             = "gpt-4-32k-0314"
	GPT432K                 = "gpt-4-32k"
	GPT40314                = "gpt-4-0314"
	GPT4                    = "gpt-4"
	GPT3Dot5Turbo0301       = "gpt-3.5-turbo-0301"
	GPT3Dot5Turbo           = "gpt-3.5-turbo"
	GPT3TextDavinci003      = "text-davinci-003"
	GPT3TextDavinci002      = "text-davinci-002"
	GPT3TextCurie001        = "text-curie-001"
	GPT3TextBabbage001      = "text-babbage-001"
	GPT3TextAda001          = "text-ada-001"
	GPT3TextDavinci001      = "text-davinci-001"
	GPT3DavinciInstructBeta = "davinci-instruct-beta"
	GPT3Davinci             = "davinci"
	GPT3CurieInstructBeta   = "curie-instruct-beta"
	GPT3Curie               = "curie"
	GPT3Ada                 = "ada"
	GPT3Babbage             = "babbage"
)

var AllModels = []string{
	GPT432K0314,
	GPT432K,
	GPT40314,
	GPT4,
	GPT3Dot5Turbo0301,
	GPT3Dot5Turbo,
	GPT3TextDavinci003,
	GPT3TextDavinci002,
	GPT3TextCurie001,
	GPT3TextBabbage001,
	GPT3TextAda001,
	GPT3TextDavinci001,
	GPT3DavinciInstructBeta,
	GPT3Davinci,
	GPT3CurieInstructBeta,
	GPT3Curie,
	GPT3Ada,
	GPT3Babbage,
}
