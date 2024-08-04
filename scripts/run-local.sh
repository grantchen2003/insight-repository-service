#!/bin/bash

cd ..

nodemon --exec "go run cmd/repository/main.go" --watch . -e go