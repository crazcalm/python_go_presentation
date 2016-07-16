import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "string_input_cgo.so")
    lib = ctypes.CDLL(path)
    user_input = ctypes.c_char_p(b"Hello")
    print("Answer is:")
    print("-------------------")
    lib.string_input(user_input)
    print("-------------------\n\n")


if __name__ == "__main__":
    main()
