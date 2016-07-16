#Notes:

##Building shared objects files with Go
    The go command to build a shared library is as follows:

    `go build -o output_file.so  -buildmode=c-shared source_file.go`

##Example:
    `go build -o hello_world.so  -buildmode=c-shared hello_world.go`

##Warning:
    You cannot have a space between the comment and the word "export"
    Example: "// export" is a no go. The header file will not be created.

##What is a shared object file?
    Add definition

##Go can build shared object files?
    As of go1.5, golang has been able to build shared object files.

##How to Run
    After building the shared object file, run the python file.

##Reflection
    (after running the file) Lets take second to obsorb what we just did. 
    We ran a python file that loaded a shared object file, which called a go
    function that printed "Hello World" to the screen. That is pretty cool!
