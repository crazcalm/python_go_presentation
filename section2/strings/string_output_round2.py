import ctypes
import os


def main():
    path = os.path.join(os.path.abspath("."), "string_output_cgo.so")
    lib = ctypes.cdll.LoadLibrary(path)
    lib.string_output.restype = ctypes.c_char_p
    print("Answer is: ")
    print("-------------------")
    answer = lib.string_output()
    print(answer)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()
