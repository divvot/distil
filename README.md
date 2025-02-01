
# Distil (IOS)

Reversed Distil SDK (IOS) protection or "x-d-token" header

## How Distil Works?

Distil currently collects the following information

* IPhone height and width
* IPhone model
* IPhone version
* Bundle ID
* Bundle Version
* Country code
* Some UUID
* Some Vendor ID
* IPs
* Jailbreak, Simulator and Debugging flags

Then a kind of checksum is made based on the previous data to solve a challenge, which is achieved with the “application_manifest”, and sent with the respective data.

---

It was actually a nice challenge for me (I'm a noob to RE).

P.S: I programmed this some time ago, I have no idea if they changed anything.