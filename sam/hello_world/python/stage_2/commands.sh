# create an s3 bucket
aws s3 mb s3://curtis-code-sam

# package cloudformation
aws cloudformation package --s3-bucket curtis-code-sam --template-file template.yaml --output-template-file gen/template-generated.yaml

# deploy cloudformation
aws cloudformation deploy --template-file gen/template-generated.yaml --stack-name sam-python-app --capabilities CAPABILITY_IAM