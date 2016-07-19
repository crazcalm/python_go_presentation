#!/bin/sh
go build -o hello_world.so  -buildmode=c-shared hello_world.go
