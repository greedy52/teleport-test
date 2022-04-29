package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	config, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	identity, err := sts.NewFromConfig(config).GetCallerIdentity(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(aws.ToString(identity.Arn))
}
