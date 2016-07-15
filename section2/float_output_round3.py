import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "float_output_cgo.so")
    lib = ctypes.CDLL(path)
    print("Answer is: ")
    print("-------------------")
    answer = ctypes.c_float(lib.float_output())
    print(answer)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()

