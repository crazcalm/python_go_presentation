import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "string_input.so")
    lib = ctypes.CDLL(path)
    user_input = ctypes.create_string_buffer(b"Hello", 10)
    print("Answer is:")
    print("-------------------")
    lib.string_input(user_input)
    print("-------------------\n\n")


if __name__ == "__main__":
    main()
