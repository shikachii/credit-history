GOOS=linux go build -o bin/main

cd bin
rm notifyCreditHistory.zip
zip notifyCreditHistory.zip main