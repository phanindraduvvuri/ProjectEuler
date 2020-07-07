#include <stdio.h>

#define MAX_SIZE 105000

int isPrime[MAX_SIZE];
void sieveOfErastosthenes() {
    for (long int i = 0; i < MAX_SIZE; i++)
    {
        isPrime[i] = 1;
    }

    for (int p = 2; p * p < MAX_SIZE; ++p)
    {
        if (isPrime[p])
        {
            for (int i = p * p; i < MAX_SIZE; i += p)
            {
                isPrime[i] = 0;
            }
        }
    }
}

int main(int argc, char const *argv[])
{
    sieveOfErastosthenes();

    int cases;
    scanf("%d", &cases);
    for (int c = 0; c < cases; ++c)
    {
        int n, cnt = 1;
        scanf("%d", &n);

        for (int i = 2; i < MAX_SIZE; ++i)
        {
            if (isPrime[i])
            {
                if (cnt == n) {
                    printf("%d\n", i);
                    break;
                }
                cnt++;
            }
        }
    }

    return 0;
}
