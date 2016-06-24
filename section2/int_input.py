import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "int_input.so")
    lib = ctypes.CDLL(path)
    user_input = int(input("Enter an integer: "))
    print("Answer is: ")
    print("-------------------")
    lib.int_input(user_input)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()

