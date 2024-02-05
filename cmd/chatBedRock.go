package cmd

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"

	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

type Request struct {
	Prompt        string   `json:"prompt"`
	MaxTokens     int      `json:"maxTokens"`
	Temperature   float64  `json:"temperature,omitempty"`
	StopSequences []string `json:"stop_sequences,omitempty"`
}

type Response struct {
	Completion string `json:"completion"`
}

type StreamingOutputHandler func(ctx context.Context, part []byte) error

func ChatBD(cmdStr, model, tokens, temperature string, cfg aws.Config) string {
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
	modelId := model
	TypeContent := "application/json"
	AcceptContent := "*/*"
	output, err := brc.InvokeModelWithResponseStream(context.Background(),
		&bedrockruntime.InvokeModelWithResponseStreamInput{
			Body:        payloadJson,
			ModelId:     aws.String(modelId),
			ContentType: aws.String(TypeContent),
			Accept:      aws.String(AcceptContent),
		})
	if err != nil {
		log.Fatal("Invoke Model error: ", err)
	}

	aiResponse, err := getStreamingResponse(output, func(ctx context.Context, event []byte) error {
		fmt.Println(string(event))
		return nil
	})

	if err != nil {
		log.Fatal("Get Streaming Response error: ", err)
	}
	return aiResponse.Completion
}

func getStreamingResponse(output *bedrockruntime.InvokeModelWithResponseStreamOutput, handler StreamingOutputHandler) (Response, error) {

	var combinedResult string
	resp := Response{}

	for event := range output.GetStream().Events() {
		switch v := event.(type) {
		case *types.ResponseStreamMemberChunk:

			//fmt.Println("payload", string(v.Value.Bytes))

			var resp Response
			err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&resp)
			if err != nil {
				return resp, err
			}

			handler(context.Background(), []byte(resp.Completion))
			combinedResult += resp.Completion

		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)

		default:
			fmt.Println("union is nil or unknown type")
		}
	}

	resp.Completion = combinedResult

	return resp, nil
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
