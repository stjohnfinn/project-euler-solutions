#include <iostream>
#include <math.h>

bool isPrime(int x) {
    for (int i = 2; i < sqrt(x) + 1; i++) {
        if (x % i == 0) {
            return false;
        }
    }
    return true;
}

int main() {
    int primeCount = -1;
    int i = 0;
    while (true) {
        if (isPrime(i))
            primeCount++;
        if (primeCount == 10001) {
            std::cout << i;
            break;
        }
        i++;
    }

    return 0;
}
