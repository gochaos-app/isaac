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

	varCfg := GetAwsCfg()
	brc := bedrockruntime.NewFromConfig(*varCfg.AwsConfig)
	tokensInt, _ := strconv.Atoi(varCfg.InitConfig.Tokens)
	temperature64, _ := strconv.ParseFloat(varCfg.InitConfig.Temperature, 64)
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
			ModelId:     aws.String(varCfg.InitConfig.Model),
			ContentType: aws.String(TypeContent),
			Accept:      aws.String(AcceptContent),
		})
	if err != nil {
		log.Fatal("Invoke Model error: ", err)
	}
	//var resp Response
	var resp map[string]interface{}
	//text := result["completions"]
	err = json.Unmarshal(output.Body, &resp)
	if err != nil {
		log.Fatal("failed to unmarshal", err)
	}
	text := getResponse(resp)
	fmt.Println(text)
	return text
}
func getResponse(resp map[string]interface{}) string {
	var textStr string
	if completions, ok := resp["completions"]; ok {
		// Loop over the slice
		for _, completion := range completions.([]interface{}) {
			// Extract the data
			if data, ok := completion.(map[string]interface{})["data"]; ok {
				// Extract the text
				if text, ok := data.(map[string]interface{})["text"]; ok {
					textStr = text.(string)
				}
			}
		}
	}
	return textStr
}
