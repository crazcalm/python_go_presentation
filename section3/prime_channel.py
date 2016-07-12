import ctypes
import os

def main():
    path = os.path.join(os.path.abspath("."), "prime_channel.so")
    lib = ctypes.CDLL(path)
    print("Counting the Primes!")
    user_input = int(input("Enter an integer: "))
    print("Answer is: ")
    print("-------------------")
    answer =  0 if user_input < 2 else lib.num_of_primes(user_input)
    print(answer)
    print("-------------------\n\n")

if __name__ == "__main__":
    main()
