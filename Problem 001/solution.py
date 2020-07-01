# num_sum = 0

# for i in range(3, 1000):
#     if i % 5 == 0 or i % 3 == 0:
#         num_sum += i

# print(num_sum)

'''Better Solution'''

total_sum = 0


def multiples_sum(num, k):
    diff = num // k
    return k * ((diff) * (diff + 1)) // 2


for _ in range(int(input())):
    N = int(input()) - 1
    result = multiples_sum(N, 3) + multiples_sum(N, 5) - multiples_sum(N, 15)
    print(result)
