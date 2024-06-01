#!/bin/bash

cd ..

export ENV=dev

nodemon --exec go run main.go