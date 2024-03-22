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

type Response struct {
	Completions []Completion `json:"completions"`
}

type Completion struct {
	Data Data `json:"data"`
}

type Data struct {
	Text string `json:"text"`
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
	var resp Response
	if err := json.Unmarshal(output.Body, &resp); err != nil {
		log.Fatal("failed to unmarshal", err)
	}

	err = json.Unmarshal(output.Body, &resp)
	if err != nil {
		log.Fatal("failed to unmarshal", err)
	}

	text := resp.Completions[0].Data.Text
	fmt.Println(text)
	return text
}
