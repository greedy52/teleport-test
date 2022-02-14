package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const awsAppPrefix = "/Users/stevehuang/.tsh/keys/teleport.dev.aws.stevexin.me/STeve-app/teleport.dev.aws.stevexin.me/awsconsole-"

func useSharedFileAndCABundle() {
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", awsAppPrefix+"credentials")
	os.Setenv("AWS_CA_BUNDLE", awsAppPrefix+"localhost-ca.pem")
}

func useAccessIDAndSecret() {
	os.Setenv("AWS_ACCESS_KEY_ID", "454c9e4611365835")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "97f852fa42d06904609521a43e9549ff17387114")
	os.Setenv("AWS_CA_BUNDLE", awsAppPrefix+"localhost-ca.pem")
}

func main() {
	// either one works
	//useAccessIDAndSecret()
	useSharedFileAndCABundle()

	os.Setenv("AWS_REGION", "ca-central-1")
	endpoint := "https://localhost:34567"

	config, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	stsClient := sts.NewFromConfig(config, func(options *sts.Options) {
		options.EndpointResolver = sts.EndpointResolverFromURL(endpoint)
	})

	identity, err := stsClient.GetCallerIdentity(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(aws.ToString(identity.Arn))
}
