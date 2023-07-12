GOOS=linux go build -o bin/main

cd bin
rm function.zip
zip function.zip main