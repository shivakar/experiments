CC=g++
CFLAGS=-std=c++11 -g

all: singleton

singleton: singleton.cpp singleton.hpp
	$(CC) $(CFLAGS) $@.cpp -o $@

clean:
	rm -f *.o singleton
	rm -rf *.dSYM
