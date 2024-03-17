package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
)

func TextFromDoc(file, prompt string) {

	varCfg := GetAwsCfg()
	txt := textract.NewFromConfig(*varCfg.AwsConfig)

	fileExtract, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	resp, err := txt.DetectDocumentText(context.Background(),
		&textract.DetectDocumentTextInput{
			Document: &types.Document{
				Bytes: []byte(fileExtract),
			},
		})

	if err != nil {
		fmt.Println("Error extracting text", err)
	}
	var textFromFile string
	for i := 1; i < len(resp.Blocks); i++ {
		if resp.Blocks[i].BlockType == "LINE" {
			//fmt.Println(*resp.Blocks[i].Text)
			textFromFile += *resp.Blocks[i].Text + "\n"
		}
	}

	if prompt == "" {
		prompt = "This text is taken from a file, file can be document, image or a PDF file, could you please review the text and tell me a small summary, what is this about or what is the main idea of this text? \n"
	}
	fmt.Println("------------------------------------------")
	fmt.Println("Text from file: ", file)
	fmt.Println("------------------------------------------")
	fmt.Println(textFromFile)
	fmt.Println("------------------------------------------")
	fmt.Println("Summary: ")
	fmt.Println("------------------------------------------")
	getOutputFromBedrock := ChatBD(prompt + textFromFile)
	fmt.Println(getOutputFromBedrock)
	fmt.Println("------------------------------------------")

}
