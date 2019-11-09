# Black Alps

## Google Bug Hunters

**Speaker: Eduardo Vela Nava**

On average the VRP (Vulnerability Reward Program) receives 35 reports per day.

There are 12 people on rotation, 2 per day checking the reports, a junior paired with a senior.

The technical review is separated from the financial rewards.

The overview of the process is the following:

```
product team <-----> triage <-----> bug hunter
|			|		^
|			|		|
|			V		|
|__________________ VRP panel __________|
```

People in the VRP panel see only the issues that are accepted (hence the most interesting).

When the bug hunter asks for a higher reward it is a good sign because it means that the vulnerability was not well understood.

In order to avoid burnout, the job is made more interesting by making bug hunters deliver better reports, how? Adding form submission tips and asking questions (e.g. how is this vulnerability exploitable?)

There is also a vulneraribility research grant: this prevents a bug hunter to send more (invalid) reports to increase 
the probability of having one accepted. The bug hunter is payed anyhow and he has an incentive to deliver only good reports.

The success rate is 13.7% (which is could because it is almost 1337, nice joke from the speaker)

Side projects: escala8 (init.g, bugswat, CTF)

## New tales of wireless security devices

**Speaker: Gerhard Klostermeier, Matthias Deeg**

Tools:
* HackRFone
* Crazyradio PA
* Logitech unifying receiver

Attack types:
* Keystroke injection
* Replay
* Data recovery
* Mouse spoofing

Techniques:
* Hardware analysis
* Radio analysis
* Firmware analysis

Take away: choose wireless devices (mouse, keyboard, etc.) wisely

## Computing over encrypted data

**Speaker: Pascal Junod @cryptopathe**

More a list of things to take a look at after the presentation given their complexity:
* FHE (Fully Homomoprhic Encryption)
* MPC (Multiparty computation)
* RLWE (Ring learning with errors)
* Oblivious transfer

Current implementations:
* [HElib](https://github.com/shaih/HElib)
* [SEAL](https://github.com/Microsoft/SEAL)

Example of applications for MPC:
* Compute an average without revealing the single addends
* Evaluate `a>b` without disclosing `a` and `b`

Attack models: semihonest (passive) vs malicious (active)

Private intersection: compute the intersection of two sets without revealing the sets

Applications: contacts list, suspects list, compromised accounts list 

Currently implemented in Chrome.

## Build a WAF as a service and forget about false positives

**Speaker: Juan Berner @89berner**

A WAF can be deployed in different ways: 
* an appliance in front of the webapp
* an agent running on the same host of the webapp
* in the cloud

Disadvantages of a WAF:
* Network placement
* Performance
* False positives

Idea: don't send all the traffic through the WAF (inline), send it also directly to the webapp (out of band)

How to direct traffic then? "Fingerprint based routing", a combination of src IP and user agent.

IPs coming from TOR exit nodes, proxies and certain IPs are routed inline (throguh the WAF). They do not represent the average customer traffic.

Another idea is to look at patterns in requests: a blind sql requires multiple queries to be effective. After n requests that look suspicious all the following from that IP are blocked.

## How to prevent a whole class of vulnerabilities

**Speaker: Sam Lanning @samlanning**

Use variant analysis: given a vulnerability write a query that finds similar patterns to identify potential vulnerabilities of the same kind across the code base.

Next step: integrate the query in the CI process.

```
bug -> diagnose -> write query -------------> monitor
	|		|	    		|
	|		V	   		V
	| 	   discover variants 	discover unreleased variants
	|		|			|
	V		V			V
	fix		fix			fix
```

## Maintaining crypto libraries for 11 languages

**Speaker: Anastasiia Voitova @vixentael**

Stats: 17% of bugs come from crypto libs, 83% from misuse of crypto libs inside apps

There are a lot of details to consider/choose when encrypting: 
* key rotation
* key exchange
* key revocation
* key derivation
* key storage
* key management system 
* operation mode
* padding
* IV
* cipher


Goal: provide to developers a crypto library that is easy to use and hard to misuse.

[Themis](https://github.com/cossacklabs/themis)

## Verifiable delay functions

**Speaker: Antonio Sanso @asanso**

The idea is to replace proof of work algorithms (which can be parallelized and thus have a consistent impact on the energy consumption of the planet) with verifiable delay functions which serve the same purpose but can not be parallelized.

## Cyber-Security: no risks no fun? 

**Speaker: Florian Schuetz**

**Disclaimer: the keynote is not meant to be political but just an open discussion**

It is hard to explain to people why cybersecurity matters and justify government investments.

The goal of cybersecurity is to safeguard our lifestyle. 

Which countries have good cybersecurity? Russia, Estonia, Israel, US.

Why? It is a matter of geopolitics. 

Israel has a small market. To get money in they have to trade something. Start ups. Sell to who? US as there are long standing relationships and
it is an open market (way easier than selling to Europe). However start-ups should not be sold otherwise money won't flow anymore in the state.

**The takeway is that Israel turned a necessity (cybersecurity) into an advantage**

How can Switzerland do the same?

Cybersecurity ensures the continuity of financial assets, of neutrality and direct democracy.

In the last year there has been an increase of 30% in DDOS attacks (sorry I don't have the reference)

For businesses this means paying more money to the ISPs to front such threats.

An idea for the cybersecurity strategy is to build an excellent network backbone. This will lower costs for business investments.

Another idea is to have more security in the education of software developers. Interesting debate question: if an engineer is responsible for a building, shouldn't a SW developer be somehow responsible for insecure code?

