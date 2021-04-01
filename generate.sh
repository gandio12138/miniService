#!/bin/bash
# shellcheck disable=SC2034
protoFilePath=$1
protoBufFilePath=$2

# shellcheck disable=SC2012
# shellcheck disable=SC2010
files=$(ls "$protoFilePath" | grep .proto)
for fileName in $files; do
  echo "$protoFilePath$fileName"
  protoc -I. --go_out=plugins=micro:"$protoBufFilePath" "$protoFilePath$fileName"
done
