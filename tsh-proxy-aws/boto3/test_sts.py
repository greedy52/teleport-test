import boto3
import json

endpoint_url = "https://127.0.0.1:34567"
region = "ca-central-1"

sts = boto3.client('sts', endpoint_url=endpoint_url, region_name=region)
print(json.dumps(sts.get_caller_identity(), indent=2))
