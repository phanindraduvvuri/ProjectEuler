import math
# import time


def divisors(num):
    num_of_divisors = 0

    for i in range(1, math.ceil(math.sqrt(num)) + 1):
        if num % i == 0:
            num_of_divisors += 2

        if i * i == num:
            num_of_divisors -= 1

    return num_of_divisors


# start = time.time()
nthTriangle, cnt = 1, 0

num = 1
while True:
    nthTriangle = (num * (num + 1)) // 2

    if num % 2 == 0:
        cnt = divisors(num // 2) * divisors(num + 1)
    else:
        cnt = divisors(num) * divisors((num + 1) // 2)

    num += 1
    if cnt >= 500:
        print(nthTriangle)
        break


# end = time.time()
# print("Program took {}s".format(end - start))
