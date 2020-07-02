#include <stdio.h>


int gcd(int a, int b) {
    if (b == 0)
        return a;
    else
        return gcd(b, a % b);
}

long long lcm(int *arr, int size) {
    long long int result = 1, i;
    for (i = 0; i < size; i++) {
        result = result * arr[i] / gcd(result, arr[i]);
    }

    return result;
}

int main(int argc, char const *argv[]) {
    int tests;
    scanf("%d", &tests);
    for (int i = 0; i < tests; i++) {
        int n;
        scanf("%d", &n);
        int arr[n];
        for (int i = 1; i <= n; i++) {
            arr[i - 1] = i;
        }
        printf("%lld\n", lcm(arr, n));
    }
    return 0;
}
