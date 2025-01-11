#!/bin/bash
INPUT_FILE=$1
OUTPUT_FILE=$2

go run main.go -file $INPUT_FILE | jq '.map |= (map(. | tostring | gsub("\\s+"; "")))' | sed 's/\"\[/\[/; s/\]\"/]/;' > $OUTPUT_FILE

