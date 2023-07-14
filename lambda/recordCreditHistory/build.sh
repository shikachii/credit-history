GOOS=linux go build -o bin/main

cd bin
rm recordCreditHistory.zip
zip recordCreditHistory.zip main