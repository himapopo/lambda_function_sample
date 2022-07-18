#!/bin/sh
GOOS=linux go build -o bin/main main.go
zip function.zip bin/main