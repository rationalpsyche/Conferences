# OWASP AppSec

## Charlie Miller - bug hunting

Choose the application to audit:
* How spread is the app?
* How hard is to find a vulnerability?
* How hard is to exploit the vulnerability?
	* Size of data: eg. sms vs pdf
	* Interaction: eg. javascript vs jpeg
	* Services: eg. Netbios vs HTTP

Static analysis - cons:
* hard to measure progress, it does not scale
* problem of understanding the whole program
* `ps->free = node` instead of `active`
This bug is difficult to find with static analysis or even by a human being, as a matter of fact it was found with a fuzzer

Dynamic analysis - aka fuzzing:
* It scales
* Cons: when do you turn off the fuzzer?

Idea:
* Start with fuzzing, this gives you code coverage
* Then perform static analysis
* Repeat

## Daniel Cornell - ZAP

Tool: ZAP + threadfix

Idea: feed to ZAP foreknowledge of the application so that the crawler has more info as problems are:
* hidden landing pages
* multi processes that crawlers do not traverse
* unknow/debug parameters

## Giancarlo Pellegrino - compression bombs

Compression may be mandated by protocol specifications (eg. HTTP response) or a custom feature (HTTP request)

When dealing with compression:
* Intensive task
* Data amplification
* Unbalanced client-server scenario

Correct way to validate: decompress the message size

## Felix Leder: bug hunting on the dark side

* Conficker: mis-rand for ip range
* Stuxnet: OR instead of AND
* Spam: quoted google translate
* Storm: hardcoded user agent with typo 'windowss' -> easy IDS rule
* Simple locker: hard coded key
* Zeus: forgot to set a root password for sys admin
* Dendroid: wrong use of prepared statement

## Tom Van Goethem - Timing attacks

Types:
* Direct timing: login request to server
* Crosstiming attacks: the attacker tricks the victim to do the timing
Example: if the user is not logged in or does not belong to a group

Browser based timings - Targets:
* parse (CPU process time)
* retrieve from cache (disk read time)
* store in cache (disk write time)

Limitations: network irregularities, gzip compression, rate limiting

### Links to video and slides: 
[AppSec.eu](http://2016.appsec.eu/?page_id=914)