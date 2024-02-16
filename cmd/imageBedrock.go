package cmd

import (
	"bytes"
	"context"
	"image/png"
	"log"
	"os"

	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type TitanImageRequest struct {
	TaskType              string                `json:"taskType"`
	TextToImageParams     TextToImageParams     `json:"textToImageParams"`
	ImageGenerationConfig ImageGenerationConfig `json:"imageGenerationConfig"`
}
type TextToImageParams struct {
	Text string `json:"text"`
}
type ImageGenerationConfig struct {
	NumberOfImages int     `json:"numberOfImages"`
	Quality        string  `json:"quality"`
	CfgScale       float64 `json:"cfgScale"`
	Height         int     `json:"height"`
	Width          int     `json:"width"`
	Seed           int64   `json:"seed"`
}

type TitanImageResponse struct {
	Images []string `json:"images"`
}

func ImageBD(cmdStr string) string {

	_, model, _, _, _, cfg := GetAwsCfg()
	brc := bedrockruntime.NewFromConfig(cfg)
	seed := int64(123)
	payload := TitanImageRequest{
		TaskType: "TEXT_IMAGE",
		TextToImageParams: TextToImageParams{
			Text: cmdStr,
		},
		ImageGenerationConfig: ImageGenerationConfig{
			NumberOfImages: 1,
			Quality:        "standard",
			CfgScale:       8.0,
			Height:         512,
			Width:          512,
			Seed:           seed,
		},
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	TypeContent := "application/json"
	output, err := brc.InvokeModel(context.Background(),
		&bedrockruntime.InvokeModelInput{
			Body:        payloadJson,
			ModelId:     aws.String(model),
			ContentType: aws.String(TypeContent),
		})
	if err != nil {
		log.Fatal("Invoke Model error: ", err)
	}
	var result TitanImageResponse

	err = json.Unmarshal(output.Body, &result)
	if err != nil {
		log.Fatal("failed to unmarshal", err)
	}
	aiImage := result.Images[0]
	unbased, err := base64.StdEncoding.DecodeString(aiImage)
	if err != nil {
		log.Fatal("failed to decode", err)
	}
	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		log.Fatal("bad image: failed to decode", err)
	}
	f, err := os.OpenFile("image.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("failed to open file", err)
	}
	png.Encode(f, im)
	return "Image generated and saved as image.png"
}
