# hashGo

<p align="center">
<img src="https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff">
<img src="https://img.shields.io/github/go-mod/go-version/zeropio/hashGo">
<img src="https://badges.frapsoft.com/os/v1/open-source.png?v=103">
</p>

![Image](/img/hashgo.gif)

---

# Installation

To build it we need Go. Then:
```
git clone https://github.com/zeropio/hashGo
cd hashGo; go build .
```

Or download it from [releases](https://github.com/zeropio/hashGo/releases).

And then move the binary to any path you use. 

---

# Usage

Pass a hash, a wordlist and select the mode:
```
hashgo -f <HASH> -w <WORDLIST PATH> -t <HASH TYPE>
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
