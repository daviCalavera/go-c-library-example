# setting compiler:
CC=gcc

# Object files:
objects = keyboard.o

# start default rule:
# Options:
# 	dynamic: on Mac OS
#	shared: on Linux
all:	dynamic

dynamic:
	rm -f libkeyboard.dylib
	rm -f ../libkeyboard.dylib
	$(CC) -c -fPIC keyboard.c
	$(CC) -dynamiclib -lncurses -o libkeyboard.dylib $(objects)
	cp libkeyboard.dylib ../libkeyboard.dylib
	make clean

shared:
	rm -f libkeyboard.so
	rm -f ../libkeyboard.so
	$(CC) -c -fPIC keyboard.c
	$(CC) -shared -nostartfiles -lncurses -Wl,-soname,libkeyboard.so -o libkeyboard.so $(objects)
	cp libkeyboard.so ../libkeyboard.so
	make clean

clean:
	rm -rf *.o

.PHONY: clean shared dynamic

#Now make sure that make rebuilds files if included headers change:
keyboard.o: keyboard.h
