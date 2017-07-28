#include <iostream>
#include <string>

#include "singleton.hpp"

int main() {
    //// Uncommenting the following lines of commented code should result in
    //// compilation errors
    // Singleton a;
    // Singleton* b = new Singleton();
    // Singleton c = a;

    Singleton& s = Singleton::GetInstance();

    Singleton& p = Singleton::GetInstance();

    // Checking whether the addresses of two variables are same or not
    std::string comp = (&s == &p ? "equal" : "not equal");

    std::cout << "Address of variables 's' and 'p' are " << comp << std::endl;

    return 0;
}
