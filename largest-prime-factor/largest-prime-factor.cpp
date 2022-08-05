#include <iostream>
#include <math.h>

// for every factor up to sqrt(X)
//
// check if prime
//
// if yes: set to max
//
// continue

bool isPrime(int x) {
    int x_sqrt = (int)sqrt(x) + 1;

    for (int i = 2; i < x_sqrt; i++) {
        if (x % i == 0) {
            return false;
        }
    }

    return true;
}

int main() {

    int LPF = 0;

    long int x = 600851475143;

    for (int i = 1; i < (int)sqrt(x) + 1; i++) {
        if (x % i == 0 && isPrime(i)) {
            LPF = i;
        }
    }

    std::cout << "Primality: " << (isPrime(LPF) ? "Yes" : "No") << std::endl;

    std::cout << "LPF Squared: " << LPF * LPF << std::endl;

    std::cout << LPF;

    return 0;
}
