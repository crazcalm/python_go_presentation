import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "int_output.so")
    lib = ctypes.CDLL(path)
    print("Answer is: ")
    print("-------------------")
    answer = lib.int_output()
    print(answer)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()
