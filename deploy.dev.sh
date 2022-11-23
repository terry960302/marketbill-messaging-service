rm -rf ./build/*

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
zip main.zip main

mv main ./build/
mv main.zip ./build/