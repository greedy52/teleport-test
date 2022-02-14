import boto3
import json

endpoint_url = "https://127.0.0.1:34567"
region = "ca-central-1"
table = "steve-dynamo"

dynamodb = boto3.resource('dynamodb',endpoint_url=endpoint_url, region_name=region)
table = dynamodb.Table(table)
print(json.dumps(table.scan(), indent=2, default=str))
