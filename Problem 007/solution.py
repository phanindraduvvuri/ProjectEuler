'''
def is_prime(n):
    if n < 2:
        return False
    else:
        for i in range(2, int(n ** 0.5) + 1):
            if n % i == 0:
                return False
        return True


def nth_prime(n):
    number_of_primes = 1
    if n == 1:
        return 2

    num = 3
    nth_prime = 2
    while number_of_primes < n:
        if is_prime(num):
            number_of_primes += 1
            nth_prime = num

        num += 2

    return nth_prime
'''


class NthPrime:
    def __init__(self):
        self.MAX_SIZE = 105000
        self.primes = []
        self.sieve_of_eratosthenes()

    def sieve_of_eratosthenes(self):
        is_prime = [True for _ in range(self.MAX_SIZE)]

        for p in range(2, int(self.MAX_SIZE ** 0.5)):
            if is_prime[p]:
                for i in range(p * p, self.MAX_SIZE, p):
                    is_prime[i] = False

        for p in range(2, self.MAX_SIZE):
            if is_prime[p]:
                self.primes.append(p)

    def nth_prime(self, n):
        return self.primes[n - 1]


nthprime = NthPrime()
for _ in range(int(input())):
    print(nthprime.nth_prime(int(input())))
