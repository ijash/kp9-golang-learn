#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <main_file>"
    exit 1
fi

main_file="$1"
target_dir="./build"

# Create the target directory if it doesn't exist
mkdir -p "$target_dir"

# Extract the filename without extension from the provided path
filename=$(basename -- "$main_file")
filename_noext="${filename%.*}"

# Build the Go code and place the binary in the target directory
go build -o "$target_dir/$filename_noext" "$main_file"
