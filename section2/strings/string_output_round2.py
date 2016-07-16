import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "string_output_cgo.so")
    lib = ctypes.CDLL(path)
    print("Answer is: ")
    print("-------------------")
    answer = lib.string_output()
    print(ctypes.c_char_p(answer).value)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()
