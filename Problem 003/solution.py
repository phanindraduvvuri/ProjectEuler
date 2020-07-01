import math

for _ in range(int(input())):
    n = int(input())

    largest_prime = 2
    while n % 2 == 0:
        n //= 2

    for i in range(3, int(math.sqrt(n)) + 1, 2):
        while n % i == 0:
            if i > largest_prime:
                largest_prime = i
            n //= i

    if n > 2:
        largest_prime = n

    print(int(largest_prime))
