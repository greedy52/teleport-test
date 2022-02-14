#!/bin/bash

# https://github.com/weaveworks/eksctl/blob/b4d4b63e821feee0d86d600e65e850a4340c570c/pkg/eks/api.go
endpoint="https://127.0.0.1:34567"
export AWS_SHARED_CREDENTIALS_FILE=/Users/stevehuang/.tsh/keys/teleport.dev.aws.stevexin.me/STeve-app/teleport.dev.aws.stevexin.me/awsconsole-credentials
export AWS_CLOUDFORMATION_ENDPOINT=${endpoint}
export AWS_EKS_ENDPOINT=${endpoint}
export AWS_EC2_ENDPOINT=${endpoint}
export AWS_ELB_ENDPOINT=${endpoint}
export AWS_ELBV2_ENDPOINT=${endpoint}
export AWS_STS_ENDPOINT=${endpoint}
export AWS_IAM_ENDPOINT=${endpoint}
export AWS_CLOUDTRAIL_ENDPOINT=${endpoint}
export AWS_REGION="ca-central-1"
eksctl get clusters
