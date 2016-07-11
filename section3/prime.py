import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "prime.so")
    lib = ctypes.CDLL(path)
    print("Prime Checker:")
    user_input = int(input("Enter an integer: "))
    print("Answer is: ")
    print("-------------------")
    answer = True if lib.prime(user_input) else False
    print(answer)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()
