import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "dict_input.so")
    lib = ctypes.CDLL(path)
    user_input = {"fname": "Marcus", "lname": "Willock"}
    print("Using this dict: {}".format(user_input))
    # create a dict
    lib.dict_input(user_input)


if __name__ == "__main__":
    main()
