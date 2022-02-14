#!/bin/bash

export NODE_EXTRA_CA_CERTS=/Users/stevehuang/.tsh/keys/teleport.dev.aws.stevexin.me/STeve-app/teleport.dev.aws.stevexin.me/awsconsole-localhost-ca.pem
export DYNAMO_ENDPOINT=https://127.0.0.1:34567
export AWS_ACCESS_KEY_ID=454c9e4611365835
export AWS_SECRET_ACCESS_KEY=97f852fa42d06904609521a43e9549ff17387114
export AWS_REGION="ca-central-1"
dynamodb-admin
