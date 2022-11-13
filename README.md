# hashGo

<p align="center">
<img src="https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff">
<img src="https://img.shields.io/github/go-mod/go-version/zeropio/hashGo">
<img src="https://badges.frapsoft.com/os/v1/open-source.png?v=103">
</p>

![Image](/img/hashgo.gif)

---

# Installation

In order to install it and build it we need Go:
```
git clone https://github.com/zeropio/hashGo
cd hashGo; go build .
```

And then move the binary to some path you use.

Or download it from [releases](https://github.com/zeropio/hashGo/releases).

---

# Usage

To use it we need to pass a file with the hash (or hashes), a wordlist and select the mode:
```
hashgo -f <HASH FILE> -w <WORDLIST> -h <HASH TYPE>
```

For now, these are the supported hashes:

| **Hash** | **Flag** |
| -------- | -------- |
| MD5 | `-h md5` |
| MD4 | `-h md4` |
| NTLM | `-h ntlm` |
| SHA1 | `-h sha1` |

---

# Enjoy it!
