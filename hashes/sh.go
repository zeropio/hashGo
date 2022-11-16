package cracking

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
    "strings"

	"github.com/labstack/gommon/color"
)

func ToSH1(hash string, wordlist io.Reader, verbose bool) {

    // Hash each wordlist line and compare it with the hash
    readFile := bufio.NewScanner(wordlist)
    readFile.Split(bufio.ScanLines)
    for readFile.Scan() {
        line := readFile.Text()
        hasher := sha1.New()
        hasher.Write([]byte(line))
        hashString := hex.EncodeToString(hasher.Sum(nil))
        pass := string(hashString[:])

        if (pass == strings.ToLower(hash)) {
            fmt.Printf(color.Cyan("The password is: ") + color.Green(line) + "\n")
        } else {
            if verbose {
                fmt.Printf(color.Red("[-]") + color.Red(line) + "\n")
            }
            continue
        }
        break
    }
    fmt.Printf(color.Cyan("Cracking for ") + color.Yellow(hash) + color.Cyan(" ended\n"))
}
