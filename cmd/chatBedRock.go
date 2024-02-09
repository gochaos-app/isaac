package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type Request struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"maxTokens"`
	Temperature float64 `json:"temperature,omitempty"`

	StopSequences []string `json:"stop_sequences,omitempty"`
}

func ChatBD(cmdStr string) string {

	model, tokens, temperature, _, cfg := GetAwsCfg()
	brc := bedrockruntime.NewFromConfig(cfg)
	tokensInt, _ := strconv.Atoi(tokens)
	temperature64, _ := strconv.ParseFloat(temperature, 64)
	payload := Request{
		Prompt:      cmdStr,
		MaxTokens:   tokensInt,
		Temperature: temperature64,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	TypeContent := "application/json"
	AcceptContent := "*/*"
	output, err := brc.InvokeModel(context.Background(),
		&bedrockruntime.InvokeModelInput{
			Body:        payloadJson,
			ModelId:     aws.String(model),
			ContentType: aws.String(TypeContent),
			Accept:      aws.String(AcceptContent),
		})
	if err != nil {
		log.Fatal("Invoke Model error: ", err)
	}
	var result map[string]interface{}

	err = json.Unmarshal(output.Body, &result)
	if err != nil {
		log.Fatal("failed to unmarshal", err)
	}
	aiResponse := getResponse(result)

	fmt.Println(aiResponse)
	return aiResponse
}

func getResponse(resp map[string]interface{}) string {
	var textStr string
	if completions, ok := resp["completions"]; ok {
		// Assert that completions is a slice
		if completionsSlice, ok := completions.([]interface{}); ok {
			// Loop over the slice
			for _, completion := range completionsSlice {
				// Assert that completion is a map
				if completionMap, ok := completion.(map[string]interface{}); ok {
					// Extract the data
					if data, ok := completionMap["data"]; ok {
						// Assert that data is a map
						if dataMap, ok := data.(map[string]interface{}); ok {
							// Extract the text
							if text, ok := dataMap["text"]; ok {
								// Assert that text is a string
								if textValue, ok := text.(string); ok {
									textStr = textValue
								}
							}
						}
					}
				}
			}
		}
	}
	return textStr
}
