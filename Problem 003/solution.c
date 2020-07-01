#include <math.h>
#include <stdio.h>


long long maxPrimeFactors(long long n)
{
	long long maxPrime = -1;

	while (n % 2 == 0) {
		maxPrime = 2;
		n >>= 1; // equivalent to n /= 2
	}

	for (int i = 3; i <= sqrt(n); i += 2) {
		while (n % i == 0) {
			if (i > maxPrime)
                maxPrime = i;
			n = n / i;
		}
	}

	if (n > 2)
		maxPrime = n;

	return maxPrime;
}

int main()
{
	int t;
    scanf("%d", &t);
    for (size_t _ = 0; _ < t; _++) {
        long long n;
        scanf("%lld", &n);
        printf("%lld\n", maxPrimeFactors(n));
    }

	return 0;
}
