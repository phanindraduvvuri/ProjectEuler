def diagnol_prod(grid):
    max_prod = -1
    for i in range(len(grid) - 3):
        for j in range(len(grid[i]) - 3):
            prod = grid[i][j] * grid[i + 1][j + 1] * \
                grid[i + 2][j + 2] * grid[i + 3][j + 3]

            max_prod = max(prod, max_prod)

    return max_prod


def rev_diagnal_prod(grid):
    max_prod = -1
    for i in range(3, len(grid) - 3):
        for j in range(len(grid) - 3):
            prod = grid[i][j] * grid[i + 1][j - 1] * \
                grid[i + 2][j - 2] * grid[i + 3][j - 3]

            max_prod = max(prod, max_prod)

    return max_prod


def right_prod(grid):
    max_prod = -1
    for i in range(len(grid) - 3):
        prod = grid[i][i] * grid[i][i + 1] * grid[i][i + 2] * grid[i][i + 3]

        max_prod = max(max_prod, prod)

    return max_prod


def down_prod(grid):
    max_prod = -1
    for i in range(len(grid) - 3):
        prod = grid[i][i] * grid[i + 1][i] * grid[i + 2][i] * grid[i + 3][i]

        max_prod = max(max_prod, prod)

    return max_prod


grid = []

for row in range(20):
    grid.append(list(map(int, input().split())))


max_prod = max(diagnol_prod(grid), rev_diagnal_prod(
    grid), right_prod(grid), down_prod(grid))

print(max_prod)
