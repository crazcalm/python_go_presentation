import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "add_one_int.so")
    lib = ctypes.CDLL(path)
    user_input = int(input("Enter your full name seperated by ','s: "))
    # create a dict
    lib.dict_input(user_input)


if __name__ == "__main__":
    main()
