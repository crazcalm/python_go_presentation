go build -o float_input.so  -buildmode=c-shared float_input.go
go build -o float_output.so  -buildmode=c-shared float_output.go
go build -o float_output_cgo.so  -buildmode=c-shared float_output_cgo.go
