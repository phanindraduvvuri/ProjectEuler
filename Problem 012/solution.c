#include<stdio.h>
#include<math.h>
// #include<time.h>


long long int divisors(int long num) {
    int numberOfDivisors = 0;

    for (long int i = 1; i <= (long) ceil(sqrt(num)); i++) {
        if (num % i == 0)
            numberOfDivisors += 2;

        if (i * i == num)
            numberOfDivisors -= 1;
    }

    return numberOfDivisors;
}

int main(int argc, char const *argv[])
{
    // clock_t begin = clock();
    long long int nthTriangle, cnt;

    for (long int num = 1; ; ++num) {
        nthTriangle = (num * (num + 1)) / 2;

        if (num % 2 == 0)
            cnt = divisors(num / 2) * divisors(num + 1);
        else
            cnt = divisors(num) * divisors((num + 1) / 2);

        if (cnt >= 500) {
            printf("%lld\n", nthTriangle);
            break;
        }
    }
    // clock_t end = clock();
    // double timeSpent = (double) (end - begin) / CLOCKS_PER_SEC;
    // printf("Program took %lfs\n", timeSpent);

    return 0;
}
