#Wrapping Go in Python Presentation

##Description
	This talk will call various shared object files (created with Go) in Python and
    systematically check what we can and cannot do with the functions in the shared 
    object file. This check includes passing various types, such as int, float, lists, 
    and dicts, into the shared object file function, and checking what we can and 
    cannot return from the shared object function.

##Setup

###Requirements
	python3.4 or greater
    Go 1.6.2 or greater

###Installing Go
	One easy way to install GO is to use [**GVM**](https://github.com/moovweb/gvm) (Go Version Manager)

	The easiest way to install python is to use your OS's package manager or
    download python from the [**python.org website**](https://www.python.org/downloads/).
