#include <stdio.h>

int max(int num1, int num2) {
    if (num1 > num2)
        return num1;
    else
        return num2;
}

int main(int argc, char const *argv[])
{
    int cases;
    scanf("%d", &cases);
    for (int c = 0; c < cases; ++c)
    {
        int N, K;
        long int max_prod = -1;
        scanf("%d %d", &N, &K);

        char number[N];
        scanf("%s", number);

        for (int i = 0; i < N - 4; ++i)
        {
            long int prod = 1;
            for (int j = i; j < K + i; ++j)
            {
                prod *= number[j] - '0';
            }
            max_prod = max(prod, max_prod);
        }

        printf("%ld\n", max_prod);
    }
    return 0;
}
