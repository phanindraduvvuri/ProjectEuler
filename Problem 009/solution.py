def get_triplet_product(n):
    prod = -1
    for a in range(3, n // 3):
        b = (n ** 2 - 2 * a * n) // (2 * n - 2 * a)
        c = n - a - b

        if c ** 2 == a ** 2 + b ** 2:
            temp = a * b * c
            if temp > prod:
                prod = temp

    return prod


for _ in range(int(input())):
    print(get_triplet_product(int(input())))
