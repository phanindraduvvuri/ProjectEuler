#include <stdio.h>

int main(int argc, char const *argv[])
{
    long int unsigned t, n, c, a, b;
    scanf("%lu", &t);
    for (int i = 0; i < t; ++i)
    {
        long long int unsigned total_sum = 0;
        scanf("%lu", &n);
        a = 1;
        b = 2;

        total_sum = 0;
        while (a < n) {
            c = a + b;
            if (a % 2 == 0)
            {
                total_sum += a;
            }
            a = b;
            b = c;
        }
        printf("%llu\n", total_sum);
    }
    return 0;
}
