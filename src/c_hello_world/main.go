package main

/*
#include <stdio.h>

void sayHello() {
	printf("Hello, World!\n");
}
*/
import "C"

func main() {
	C.sayHello()
}
