terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

data "aws_region" "current" {}

provider "aws" {
  region = "ca-central-1"
  endpoints {
    dynamodb = "https://localhost:34567"
    sts      = "https://localhost:34567"
    ec2      = "https://localhost:34567"
  }
}

data "aws_caller_identity" "current" {}

output "identity" {
  value = data.aws_caller_identity.current
}

data "aws_dynamodb_table" "steve" {
  name = "steve-dynamo"
}

output "dynamo" {
  value = data.aws_dynamodb_table.steve
}
