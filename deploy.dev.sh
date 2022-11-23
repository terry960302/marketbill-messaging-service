rm -rf ./build/*

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
zip main.zip main

mv main ./build/
mv main.zip ./build/

NAME=messaging-dev
FILE_PATH=build/main.zip
aws lambda update-function-code --function-name $NAME --zip-file fileb://$FILE_PATH
