import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "float_input.so")
    lib = ctypes.CDLL(path)
    user_input = float(input("Enter a float: "))
    print("Answer is:")
    print("--------------------")
    lib.float_input(ctypes.c_double(user_input))
    print("--------------------/n/n")


if __name__ == "__main__":
    main()
