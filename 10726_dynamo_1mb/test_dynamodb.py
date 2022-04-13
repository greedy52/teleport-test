import boto3
import json

dynamodb = boto3.resource('dynamodb',region_name="ca-central-1")
table = dynamodb.Table("steve-dynamo")
conditions= {
        "name": {
            "AttributeValueList": ["STeve"],
            "ComparisonOperator": "EQ",
        },
}
filters={
        "career": {
            "AttributeValueList": ["robot"],
            "ComparisonOperator": "EQ",
        },
}

# ---- ScannedCount 4, Count 2
print(json.dumps(table.query(KeyConditions=conditions, QueryFilter=filters), indent=2, default=str))

# ---- ScannedCount 2, Count 0
print(json.dumps(table.query(KeyConditions=conditions, QueryFilter=filters, Limit=2), indent=2, default=str))

# ---- ScannedCount 3, Count 1
print(json.dumps(table.query(KeyConditions=conditions, QueryFilter=filters, Limit=3), indent=2, default=str))

# ---- ScannedCount 4, Count 2
print(json.dumps(table.query(KeyConditions=conditions, QueryFilter=filters, Limit=4), indent=2, default=str))
