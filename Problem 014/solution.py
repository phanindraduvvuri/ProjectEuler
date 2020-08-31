def collatzSequenceUtil(n, collLenDict):
    if n in collLenDict:
        return collLenDict[n]

    if n == 1:
        collLenDict[n] = 1

    elif n % 2 == 0:
        collLenDict[n] = 1 + collatzSequenceUtil(n // 2, collLenDict)
    else:
        collLenDict[n] = 1 + collatzSequenceUtil((3 * n) + 1, collLenDict)

    return collLenDict[n]


def collatzSequence(n):
    lenght = -1
    num = 0
    collLenDict = {}

    collatzSequenceUtil(n, collLenDict)

    for i in range(1, n):
        if i not in collLenDict:
            collatzSequenceUtil(i, collLenDict)

        currentLen = collLenDict[i]

        if lenght < currentLen:
            lenght = currentLen
            num = i

    return num, lenght


print(collatzSequence(1000000))
