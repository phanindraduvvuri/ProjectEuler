def sum_of_squares(n):
    return (n * (n + 1) * (2 * n + 1)) // 6


def sum_of_numbers(n):
    return (n * (n + 1)) // 2


for _ in range(int(input())):
    n = int(input())

    result = (sum_of_numbers(n) ** 2) - sum_of_squares(n)
    print(result)
