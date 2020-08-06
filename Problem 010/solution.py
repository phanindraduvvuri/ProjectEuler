import math

"""
def sum_of_primes(n):
    sieve = [False for _ in range(n + 1)]

    for i in range(4, n + 1, 2):
        sieve[i] = True

    for i in range(3, int(math.sqrt(n)) + 1, 2):
        if not sieve[i]:
            for j in range(i * i, n + 1, 2 * i):
                sieve[j] = True

    prime_sum = 0
    for i in range(2, n + 1):
        if not sieve[i]:
            prime_sum += i

    return prime_sum
"""

limit = 10 ** 6
sieve = [False for _ in range(limit + 1)]


def sieve_of_eratosthenes():
    limit = 10 ** 6
    sieve = [False for _ in range(limit + 1)]

    for i in range(4, limit + 1, 2):
        sieve[i] = True

    for i in range(3, int(math.sqrt(limit)) + 1, 2):
        if not sieve[i]:
            for j in range(i * i, limit + 1, 2 * i):
                sieve[j] = True

    sum_of_primes = [0, 1]
    sum_sofar = 0
    for i in range(2, limit + 1):
        if not sieve[i]:
            sum_sofar += i

        sum_of_primes.append(sum_sofar)

    return sum_of_primes


sum_of_primes = sieve_of_eratosthenes()
for _ in range(int(input())):

    print(sum_of_primes[int(input())])
