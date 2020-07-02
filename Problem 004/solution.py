def is_palindrom(num):
    reversed_num = 0
    number = num

    while num > 0:
        reversed_num = reversed_num * 10 + num % 10
        num //= 10

    return reversed_num == number


for _ in range(int(input())):
    n = int(input())
    largest_prime = -1

    for i in range(990, 99, -11):
        for j in range(999, 99, -1):
            prod = i * j
            if prod > largest_prime and is_palindrom(prod) and prod < n:
                largest_prime = prod
                break
            elif prod < largest_prime:
                break

    print(largest_prime)
