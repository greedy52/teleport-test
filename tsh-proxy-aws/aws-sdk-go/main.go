package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	identity, err := sts.New(session.New()).GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		panic(err)
	}

	fmt.Println(identity)
}
