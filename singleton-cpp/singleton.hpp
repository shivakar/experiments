#include <iostream>

class Singleton {
   public:
    // GetInstance returns "the" instance of the Singleton class
    // GetInstance is guaranteed to be thread-safe in c++11 but not in c++03 or
    // below
    static Singleton& GetInstance() {
        static Singleton s;
        return s;
    }
    ~Singleton() {
        std::cout << "Destroying the Singleton object" << std::endl;
    }
    // Disallow copying by explicitly setting to delete
    // requires std=c++11
    // If using a c++03 compiler just make declare them without implementing
    // them, which should cause compile-time warning if such a call to copy
    // constructor or assignment operator is ever made
    Singleton(Singleton&) = delete;
    Singleton& operator=(const Singleton&) = delete;

   private:
    // Disallow instantiation by external code
    Singleton() { std::cout << "Creating new Singleton object" << std::endl; }
};
