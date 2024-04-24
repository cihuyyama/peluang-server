package util

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

func SendSimpleEmail(Message string, Subject string, Recipient string) error {
	Sender := "Peluang Teams <official@peluang.co>"
	Region := "ap-southeast-1"

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(Region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return err
	}
	client := sesv2.NewFromConfig(cfg)
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &Sender,
		Destination: &types.Destination{
			ToAddresses: []string{Recipient},
		},
		ReplyToAddresses: []string{Sender},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &Message,
					},
				},
				Subject: &types.Content{
					Data: &Subject,
				},
			},
		},
	}

	res, err := client.SendEmail(ctx, input)
	if err != nil {
		log.Fatalf("Error sending email, %v", err)
		return err
	}
	fmt.Println(res)
	return nil
}

func SendTemplatedEmailVerification(otp int64, Recipient string) error {
	Sender := "Peluang Teams <official@peluang.co>"
	ReplyTo := "official@peluang.co"
	Region := "ap-southeast-1"
	OTP := fmt.Sprintf("{\"OTP\":\"%d\"}", otp)
	TempName := "PeluangOTPTemplate"

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(Region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return err
	}
	client := sesv2.NewFromConfig(cfg)
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &Sender,
		Destination: &types.Destination{
			ToAddresses: []string{Recipient},
		},
		ReplyToAddresses: []string{ReplyTo},
		Content: &types.EmailContent{
			Template: &types.Template{
				TemplateName: &TempName,
				TemplateData: &OTP,
			},
		},
	}

	res, err := client.SendEmail(ctx, input)
	if err != nil {
		log.Fatalf("Error sending email, %v", err)
		return err
	}
	fmt.Println(res.MessageId)
	return nil
}

func CreateTemplateEmail(TempSubject string, TempName string, TempHtml string) error {
	Region := "ap-southeast-1"

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(Region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return err
	}
	client := sesv2.NewFromConfig(cfg)

	template := &sesv2.CreateEmailTemplateInput{
		TemplateName: &TempName,
		TemplateContent: &types.EmailTemplateContent{
			Subject: &TempSubject,
			Html:    &TempHtml,
		},
	}

	_, err = client.CreateEmailTemplate(ctx, template)
	if err != nil {
		log.Fatalf("Error creating template, %v", err)
		return err
	}
	return nil
}
