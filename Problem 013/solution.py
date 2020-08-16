sum_val = 0


with open('input') as f:
    # sum_val += (f.readline())
    while True:
        line = f.readline().strip()

        if not line:
            break
        else:
            line = int(line)

        sum_val += line

print(str(sum_val)[:10])
