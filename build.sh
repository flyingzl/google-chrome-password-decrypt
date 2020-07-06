#!/bin/bash

rm -rf output
mkdir -p output
go build -ldflags "-s -w" -o output/hackChrome_temp
upx -9 output/hackChrome_temp -o output/hackChrome
rm output/hackChrome_temp