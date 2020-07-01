#include <stdio.h>

int multiples_sum(int, int);

int main(int argc, char const *argv[])
{
    int t;
    scanf("%d",&t);
    for (int i = 0; i < t; ++i)
    {
        int N, result;
        scanf("%d", &N);
        result = multiples_sum(N - 1, 3) + multiples_sum(N - 1, 5) - multiples_sum(N - 1, 15);
        printf("%d\n", result);
    }
    return 0;
}


int multiples_sum(int num, int k)
{
    int common_diff = num / k;
    return(k * ((common_diff) * (common_diff + 1)) / 2);
}
