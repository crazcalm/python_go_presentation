#!/bin/sh
go build -o prime.so  -buildmode=c-shared prime.go
go build -o prime_channel.so  -buildmode=c-shared prime_channel.go
go build -o int_inputs.so  -buildmode=c-shared int_inputs.go
