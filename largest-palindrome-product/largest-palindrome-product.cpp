#include <iostream>
#include <string>
#include <algorithm>

bool isPalindrome(int x) {
    std::string s = std::to_string(x);
    std::string temp = s;
    reverse(s.begin(), s.end());
    return (temp == s);
}

int main() {

    int max = 0;

    for (int i = 100; i < 1000; i++) {
        for (int j = i; j < 1000; j++) {
            if (i * j > max) {
                if (isPalindrome(i * j)) {
                    max = i * j;
                }
            }
        }
    }

    std::cout << max;

    return 0;
}
