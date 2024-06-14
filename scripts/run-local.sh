#!/bin/bash

cd ..

export ENV=dev

nodemon --exec "go run cmd/repository/main.go" --watch . -e go