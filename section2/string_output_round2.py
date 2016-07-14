import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "string_output.so")
    lib = ctypes.CDLL(path)
    print("Answer is: ")
    print("-------------------")
    answer = lib.string_output()
    print(ctypes.c_wchar_p(answer))
    print("-------------------\n\n")

if __name__ == "__main__":
    main()
