#!/bin/bash

GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ~/bin-weaming/itree-linux .
go build -ldflags "-s -w" -o ~/bin-weaming/itree-macos .
