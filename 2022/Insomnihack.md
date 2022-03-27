# Insomnihack
## It’s Raining Shells - How to Find New Attack Primitives in Azure?
**Speaker.** Andy Robbins `@_wald0`

### Abuses vs Bugs

Advantages of abuses:
- cheaper to maintain
- longer shelf life
- they exist in each instance
- they are a challenge for detection

### Research methodology
- begin with the end in mind
- study design system ("academic knowledge")
	- official documentation
	- repos and repos issues
	- talks (12-24 hours)
	- chats, forums
- explore the system

OSINT example to find people to learn from:
`site:linkedin.com "microsoft" "graph" "architect"`

**Remark.** Documentation is not accurate and up to date. Go beyond documentation.

How to explore Azure

- azure GUI + developer tools
- az tool
- az powershell module

**Remark.** Not all the authorization is stored in the JWT -> probably you have more permission than what the JWT describes.

Steps:
- make a list of app roles
- test the app with one permission at time
- automate the process

**Resources.**
- Coming up: Atomic Azure Tests -> inspired by [Atomic Red Team](https://github.com/redcanaryco/atomic-red-team/)

## Automatically extracting static anti-virus signatures
**Speaker.** Vladimir Meier `@plowsec`

The AV performs multiple steps:
- check the hash
- heuristic
- signatures
- emulation

A CI/CD pipeline was set up to automate AV scans. Two reasons for this against using Virus Total:
- VT APIs are expensive
- VT results are shared with AV vendors

Heavy usage of taviso's [loadlibrary](https://github.com/taviso/loadlibrary) and `vmrun` to run an executable "headless" with VMWare.

If file is no longer present then it was detected.

One way to detect AV signatures is to perform divide et impera, see `dsplit` , and check which chunk is marked as detected.

**Issue.** plain divide et impera would lead to a corrupted PE, you need PE format aware mutations.

**Issue.** When the signature spans between two chunks is not detected correctly with divide et impera -> see interval trees

Strings can be identified with the same logic and replaced, meaning
- enumerate detected strings
- for each string enumerate the cross references
- for each cross reference inject a function that encrypts and decrypts the string. 

This is accomplished with
- LIEF
- radare2
- keystone

**Resources.**
* [avcleaner](https://github.com/scrt/avcleaner)
* [Engineering AV evasion](https://blog.scrt.ch/2020/06/19/engineering-antivirus-evasion/)

## REW-sploit: dissect payloads with ease
**Speaker.** Cesare Pizzi `@red5heep`

See [REW-sploit](https://github.com/REW-sploit/REW-sploit)

**Resource.** (unrelated) [asciicinema](https://asciinema.org/)

## **Managing large-scale response**  
**Speaker.**  Mathias Fuchs  `@mathias_fuchs`

What you need: 
- visibility
- efficiency
	- e.g. you can not always get full disk images
- technical skills
- documentation
- soft skills

### Defender tools
Always on:
- sysmon
- logs
- netflow
On demand:
- EDR
- pcap
- forensics agents
- enriched alerts

On efficiency: have a small dataset from many machines to understand where to focus and then have a large dataset from few interesting machines.

Interesting finding: capacity management often detects ransomware earlier than AV/EDR. For instance a NAS is full

Detection can be done by counting number of occurrences of certain events, outliers will point out systems to check.

As incident responder you want to minimize false negatives: e.g. when matching IOCs you might have a hash (md5) but the tool only supports  sha1,sha256 and then it would not return a match.

As incident responder never do anything without a guiding question.

### Documentation
Documentation keeps control of the case: you can brief the customer at any point in time.

"spreadsheet of DOM" contains:
- master timeline and relevant IOCs
- contacts

Working with a flat file could be convenient in no internet situations (e.g. flights or air gapped systems)

**Tool.** [Aurora IR](https://www.sans.org/tools/aurora-ir/)

The timeline should show if something is missing (false negatives).

On resource management and analyst not working at night: customer needs to see progress, it does not matter when progress was made.

## A Common Bypass Pattern to Exploit Modern Web Apps
**Speaker.** Simon Scannell `@scannell_simon`

Core idea of the talk delivered via four vulnerabilities is: don't stop looking when input is sanitized, continue to see if the sanitized input is then modified.

On a high level input is:
- transformed
- normalized
- sanitized
- **potentially edited** 

Easy example to remember: file uploads checks file extension however DB truncates the file name. Provide `AAAA...AAA.php.xml`

## Hook, Line and Sinker - Pillaging API Webhooks
**Speaker.** Abhay Bhargav

A webhook is a user generated callback. 

Provider: e.g. stripe, mailchimp, CI/CD 
Consumer: developer app

The core idea of the talk is to attack the provider via SSRF upon a webhook call.

The provider sends a POST request to the consumer upon webhook. The (malicious) consumer replies with `HTTP 303`  "follow redirect" to e.g. `https://169.254.169.254/metadata`. The response is then displayed on the provider page or in the source code.

Main mitigation: disable default follow redirect in clients.