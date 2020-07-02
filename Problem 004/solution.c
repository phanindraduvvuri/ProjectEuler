#include <stdio.h>

int isPalindrom(int num) {
    int reversed = 0, number = num;

    while (num > 0) {
        reversed = reversed * 10 + num % 10;
        num /= 10;
    }

    return (reversed == number);
}

int main(int argc, char const *argv[]) {
    int test;
    scanf("%d", &test);
    for (int c = 0; c < test; c++) {
        int n;
        scanf("%d", &n);
        long long int prod, largestPrime = -1;

        for (int i = 990; i > 99; i-=11) {
            for (int j = 999; j > 99; j--) {
                prod = i * j;
                if ((prod > largestPrime) && (isPalindrom(prod)) && (prod < n)) {
                    largestPrime = prod;
                    break;
                }
                else if (prod < largestPrime)
                    break;
            }
        }
        printf("%lld\n", largestPrime);
    }
    return 0;
}
