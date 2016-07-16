import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "list_input.so")
    lib = ctypes.CDLL(path)
    user_input = [1,2,3,4]
    print("Will use {}".format(user_input))
    lib.list_input(user_input)


if __name__ == "__main__":
    main()
