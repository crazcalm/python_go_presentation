#!/bin/sh
go build -o int_input.so  -buildmode=c-shared int_input.go
go build -o int_output.so  -buildmode=c-shared int_output.go
