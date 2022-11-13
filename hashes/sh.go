package cracking

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/labstack/gommon/color"
)

func Sh1(hash string, wordlist io.Reader) {

    // Hash each wordlist line and compare it with the hash
    readFile := bufio.NewScanner(wordlist)
    readFile.Split(bufio.ScanLines)
    for readFile.Scan() {
        line := readFile.Text()
        hasher := sha1.New()
        hasher.Write([]byte(line))
        hashString := hex.EncodeToString(hasher.Sum(nil))
        pass := string(hashString[:])

        if (pass == hash) {
            fmt.Printf(color.Cyan("The password is: ") + color.Green(line) + "\n")

        }
    }
    fmt.Printf(color.Cyan("Cracking for ") + color.Yellow(hash) + color.Cyan(" ended\n"))
}
