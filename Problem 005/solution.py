def gcd(a, b):
    if b == 0:
        return a
    else:
        return gcd(b, a % b)


def lcm(lst):
    result = 1
    for ele in lst:
        result = result * ele // gcd(result, ele)

    return result


for _ in range(int(input())):
    lst = [i for i in range(1, int(input()) + 1)]

    print(lcm(lst))
