#!/bin/bash
# shellcheck disable=SC2034
protoFilePath=$1
protoBufFilePath=$2

protoc -I. --go_out=plugins=micro:"$protoBufFilePath" "$protoFilePath*.proto"