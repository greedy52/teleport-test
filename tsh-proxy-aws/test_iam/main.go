package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	session, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	iamClient := iam.New(session)
	output, err := iamClient.ListGroups(&iam.ListGroupsInput{})
	fmt.Println(output)
	fmt.Println(err)
}
