#!/bin/bash

cd ../internal/services

for dir in */ ; do
    # Check if the item is a directory
    if [ -d "$dir" ]; then
        echo "Entering directory: $dir"
        cd "$dir" || exit
        # Execute the specified command
        protoc *.proto --go_out=.. --go-grpc_out=..
        # Return to the parent directory
        cd ..
    fi
done