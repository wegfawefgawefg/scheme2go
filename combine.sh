#!/bin/bash

# Define the output file
output="./combined.go"

# Create or clear the output file
> "$output"

# List of files to concatenate
files=(
    "./pkg/compile/compile.go"
    "./pkg/emit/emit.go"
    "./pkg/parse/parse.go"
    "./pkg/tokenize/tokenize.go"
)

# Loop through the files and append each to the output file with a newline in between
for file in "${files[@]}"; do
    if [ -f "$file" ]; then
        cat "$file" >> "$output"
        echo -e "\n" >> "$output"  # Append a newline after the content of each file
    else
        echo "Warning: '$file' does not exist and will be skipped."
    fi
done

echo "Files have been combined into $output."
