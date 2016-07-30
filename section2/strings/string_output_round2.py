import ctypes
import os


def main():
    path = os.path.join(os.path.abspath("."), "string_output_cgo.so")
    lib = ctypes.cdll.LoadLibrary(path)
    lib.string_output.restype = ctypes.c_char_p
    answer = lib.string_output()
    print("Answer is: ")
    print("-------------------")
    print(answer)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()
