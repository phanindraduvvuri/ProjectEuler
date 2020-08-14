#include <stdio.h>

int gridSize = 20;

int max(int val1, int val2)
{
    if (val1 > val2)
        return val1;

    return val2;
}

long long int diaProd(int grid[][20])
{
    long long int maxProd = -1, prod;

    for (int i = 0; i < gridSize - 3; i++)
    {
        for (int j = 0; j < gridSize - 3; j++)
        {
            prod = grid[i][j] * grid[i + 1][j + 1] * grid[i + 2][j + 2] * grid[i + 3][j + 3];

            maxProd = max(maxProd, prod);
        }
    }

    return (maxProd);
}

long long int revDiaProd(int grid[][20]) {
    long long int maxProd = -1, prod;

    for (int i = 3; i < gridSize - 3; i++) {
        for (int j = 3; j < gridSize - 3; j++) {
            prod = grid[i][j] * grid[i + 1][j - 1] * grid[i + 2][j - 2] * grid[i + 3][j - 3];

            maxProd = max(maxProd, prod);
        }
    }

    return (maxProd);
}

long long int rightProd(int grid[][20]) {
    long long int maxProd = -1, prod;

    for (int i = 0; i < gridSize - 3; i++) {
        prod = grid[i][i] * grid[i][i + 1] * grid[i][i + 2] * grid[i][i + 3];

        maxProd = max(maxProd, prod);
    }

    return (maxProd);
}

long long int downProd(int grid[][20]) {
    long long int maxProd = -1, prod;

    for (int i = 0; i < gridSize - 3; i++) {
        prod = grid[i][i] * grid[i + 1][i] * grid[i + 2][i] * grid[i + 3][i];

        maxProd = max(maxProd, prod);
    }

    return (maxProd);
}

int main(int argc, char const *argv[])
{
    int grid[gridSize][gridSize];
    long long int res[4], result = 0;

    for (int i = 0; i < gridSize; ++i)
    {
        for (int j = 0; j < gridSize; ++j)
        {
            scanf("%d", &grid[i][j]);
        }
    }

    res[0] = diaProd(grid);
    res[1] = revDiaProd(grid);
    res[2] = rightProd(grid);
    res[3] = downProd(grid);
    for (int i = 0; i < 4; i++) {
        result = max(res[i], result);
    }

    printf("%lld\n", result);

    return 0;
}
