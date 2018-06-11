package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func handler(ctx context.Context, evt json.RawMessage) (string, error) {

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return "", err
	}
	cfg.Region = endpoints.UsWest2RegionID

	svc := ses.New(cfg)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []string{
				"kai.hendry@gmail.com",
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("This message body contains HTML formatting. It can, for example, contain links like this one: <a class=\"ulink\" href=\"http://docs.aws.amazon.com/ses/latest/DeveloperGuide\" target=\"_blank\">Amazon SES Developer Guide</a>."),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("This is the message body in text format."),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Test email"),
			},
		},
		Source: aws.String("dev.case@unee-t.com"),
	}

	req := svc.SendEmailRequest(input)
	result, err := req.Send()
	if err != nil {
		return "", err
	}

	fmt.Println(result)
	return "Mail sent", err

}

func main() {
	lambda.Start(handler)
}
