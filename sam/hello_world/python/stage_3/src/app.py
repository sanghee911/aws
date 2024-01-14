import json
import os

import boto3

region = os.environ.get('REGION')
dynamo = boto3.client('dynamodb')
table_name = os.environ.get('TABLE_NAME')

def response(err, res=None):
    return {
        "statusCode": '400' if err else '200',
        "body": err.message if err else json.dumps(res),
        "headers": {
            'Content-Type': 'application/json',
        }
    }


def lambda_handler(event, context):
    print(f"Loading data from table {table_name} in {region}")
    scan_result = dynamo.scan(TableName=table_name)
    return response(None, scan_result)
