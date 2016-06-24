import ctypes
import os

def main():
    # Path to shared object file
    path = os.path.join(os.path.abspath("."), "hello_world.so")

    # load shared object file
    lib = ctypes.CDLL(path)

    # Call shared object method
    lib.helloWorld()


if __name__ == "__main__":
    main()
