# Description: Deploy lambda function to AWS
aws s3 cp bin/recordCreditHistory.zip s3://credit-history-shikachii/recordCreditHistory.zip --profile creditHistory
aws lambda update-function-code --function-name recordCreditHistory \
    --s3-bucket credit-history-shikachii \
    --s3-key recordCreditHistory.zip \
    --profile creditHistory