#!/bin/sh
go build -o string_input.so  -buildmode=c-shared string_input.go
go build -o string_input_cgo.so  -buildmode=c-shared string_input_cgo.go
go build -o string_output.so  -buildmode=c-shared string_output.go
go build -o string_output_cgo.so  -buildmode=c-shared string_output_cgo.go
go build -o string_output_pointer.so  -buildmode=c-shared string_output_pointer.go
