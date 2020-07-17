#include <stdio.h>

long long get_triplet_prod(int n)
{
    long long int prod = -1, temp = -1;

    for (int a = 3; a < n / 3; a++)
    {
        int b = ((n * n) - (2 * a * n)) / ((2 * n) - (2 * a));
        int c = n - a - b;

        if ((c * c) == (a * a) + (b * b))
        {
            temp = a * b * c;
            if (temp > prod)
                prod = temp;
        }
    }

    return prod;
}

int main(int argc, char const *argv[])
{
    int cases, n;
    scanf("%d", &cases);
    for (int c = 0; c < cases; c++)
    {
        scanf("%d", &n);
        printf("%lld \n", get_triplet_prod(n));
    }

    return 0;
}
