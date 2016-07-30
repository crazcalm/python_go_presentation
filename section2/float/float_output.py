import ctypes
import os


def main():
    path = os.path.join(os.path.abspath("."), "float_output.so")
    lib = ctypes.CDLL(path)
    lib.float_output.restype = ctypes.c_double
    print("Answer is: ")
    print("-------------------")
    answer = lib.float_output()
    print(answer)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()
