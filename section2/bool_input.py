import ctypes
import os
from random import randint

def main():
    path = os.path.join(os.path.abspath("."), "bool_input.so")
    lib = ctypes.CDLL(path)
    user_input = True if randint(0,1) else False
    print("Answer is: ")
    print("-------------------")
    lib.bool_input(user_input)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()
