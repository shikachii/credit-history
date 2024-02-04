# GOOS=linux go build -o bin/main
GOOS=linux GOARCH=amd64 go build -o bin/main

cd bin
rm recordCreditHistory.zip
zip recordCreditHistory.zip main