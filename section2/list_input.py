import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "add_one_int.so")
    lib = ctypes.CDLL(path)
    user_input = int(input("Add a string delimerated by ','s: "))
    lib.list_input(user_input)


if __name__ == "__main__":
    main()
