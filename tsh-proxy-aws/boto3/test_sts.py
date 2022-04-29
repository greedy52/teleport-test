import boto3
import json

print(json.dumps(boto3.client('sts').get_caller_identity(), indent=2))
