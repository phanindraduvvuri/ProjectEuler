#include <stdio.h>
#include <math.h>

long long sumOfNumbers(long int n) {
    return (n * (n + 1)) / 2;
}

long long sumOfSquares(long int n) {
    return (n * (n + 1) * (2 * n + 1)) / 6;
}

int main(int argc, char const *argv[]) {
    long int tests;
    scanf("%ld", &tests);
    for (long int i = 0; i < tests; i++) {
        long int n;
        scanf("%ld", &n);

        long long int unsigned result = pow(sumOfNumbers(n), 2) - sumOfSquares(n);
        printf("%llu\n", result);
    }
    return 0;
}
