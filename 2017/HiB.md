# HackInBo

## Android - static analysis with Joern

Famous example `linssh2`: grep with regexp to find mathematical operators within ALLOC functions. 
The point is: can we do better?

Code property graph: it is a multigraph in the sense that edges and vectors can be labelled with more than one property.
Different abstractions:

* Abstract syntax tree (AST) - eg. used in compilers
* Control flow graph (CFG) - eg. taint analysis
* Program dependent graph

Recall that static analysis has limitations in detecting vulnerabilities related to function pointers, dynamic function calls, race conditions.

## `go get my/vulnerabilities`

Where to look for vulnerabilities in a language? Where new features lie.

Syscalls are slow therefore Go will execute everything it can while waiting for stdint (for instance).

```
for i:=0;i<=10;i++ {
	go func() {
		fmt.Printf("%d",i)
	}
}
```

it will thus print `10` at each iteration. The workaround for this is to redefine `i` within the loop.

Channels can be a cause of information leakage. Imagine a webserver written in Go. It receives two http requests, in this scenario the example above would translate in a response to the first client with the output of the second.

### Core-dependent behaviour

The architecture: 

`OS threads -> Go logic processor -> go routines`

Another issue is different behaviour with Go routines between machines with single and more processors.
You can image the infinite loop in the slides as representing computationally long tasks.

### Garbage collector

As long as all go routines do not end/retails the gargabe collector can not take action hence you could end up finishing the memory.

Another issue: it exists a deadlock error notification BUT even this does not happen if runtime is blocked.

### Solutions

Workaround: channels (good solution) or use of non-inline functions (bad solution).
The latter is discouraged for the following reasons:
* horrible benchmarks
* unreadable code (eg. recursive function instead of for)
* if the compiler changes then the whole code could break

Go routines - eg. `go foo()` - do not have return values. Solution: use `select` reading from more than one channel.
With `default` option reading does not block.

### Read sources before trusting libraries

`TimeoutHandler` is similar to `set_time_limit` in PHP. However it takes a handler written by the user so it is unclear how such function could work as expected.

When you write or review Go code check the source of the library you are using.

Last advice: always close channels

Examples in `source/`

## From APK to Golden Ticket

See [detailed article](http://resources.infosecinstitute.com/apk-golden-ticket-owning-android-smartphone-gaining-domain-admin-rights/)

Note: often virtual servers are cloned and the password is not changed!

## Brida

WebApp | Mobile App
--- | ---
Fixed client | Custom compiled
Logic on backend | Logic split between backend and frontend
Provisioned from the server | Provisioned from a trusted third party (i.e. play store)

We can break down a mobile assessment into levels:

* **Level 1:** no encryption

	http traffic analysis
* **Level 2:** TLS involved

	MITM with Burpsuite, install Burp CA
* **Level 3:** Certificate Pinning

	Use FRIDA
* **Level 4:** POST body encrypted

	Find hard-coded AES key
* **Level 5:** The key is not fixed

	The scheme is the following:

	* CT = AES(PT, random_key)
	* send(CT, RSA_enc(random_key))
	* The client holds the public RSA key.

	Hooking on the key: keep it fixed

* **Level 6:** the server checks the keys to be different

## Cuckoo sandbox

Do spend time setting a proper VM for malware analysis.

Stupid checks a malware performs to detect VM:
* 10 GB hard disk?
* Default settings? (e.g. screensaver)
* No saved passwords in browser?

**Remark.** the analysis environment is NOT trusted: malware could target the hypervisor
