package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

const awsAppPrefix = "/Users/stevehuang/.tsh/keys/teleport.dev.aws.stevexin.me/STeve-app/teleport.dev.aws.stevexin.me/awsconsole-"

func useSharedFileAndLoadConfig() {
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", awsAppPrefix+"credentials")
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true") // must be used to load ca_bundle from file
}

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
	//useSharedFileAndLoadConfig()
	useSharedFileAndCABundle()
	//useAccessIDAndSecret()

	os.Setenv("AWS_REGION", "ca-central-1")
	endpoint := "https://localhost:34567"

	session, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint))
	if err != nil {
		panic(err)
	}

	stsClient := sts.New(session, aws.NewConfig().WithEndpoint(endpoint))
	identity, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		panic(err)
	}

	fmt.Println(identity)
}
