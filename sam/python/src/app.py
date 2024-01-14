import json


def lambda_handler(event: dict, context: dict) -> dict:

    for key, value in event.items():
        print(f"{key}: {value}")

    return {
        "statusCode": 200,
        "body": json.dumps({
            "message": "hello world",
        }),
    }
