package cracking

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/labstack/gommon/color"
	"golang.org/x/crypto/md4"
)

func Md5(hash string, wordlist io.Reader) {

    // Hash each wordlist line and compare it with the hash
    readFile := bufio.NewScanner(wordlist)
    readFile.Split(bufio.ScanLines)
    for readFile.Scan() {
        line := readFile.Text()
        hasher := md5.New()
        hasher.Write([]byte(line))
        hashString := hex.EncodeToString(hasher.Sum(nil))
        pass := string(hashString[:])

        fmt.Println(pass)
        if (pass == strings.ToLower(hash)) {
            fmt.Printf(color.Cyan("The password is: ") + color.Green(line) + "\n")

        }
    }
    fmt.Printf(color.Cyan("Cracking for ") + color.Yellow(hash) + color.Cyan(" ended\n"))
}


func Md4(hash string, wordlist io.Reader) {

    // Hash each wordlist line and compare it with the hash
    readFile := bufio.NewScanner(wordlist)
    readFile.Split(bufio.ScanLines)
    for readFile.Scan() {
        line := readFile.Text()
        hasher := md4.New()
        hasher.Write([]byte(line))
        hashString := hex.EncodeToString(hasher.Sum(nil))
        pass := string(hashString[:])

        if (pass == strings.ToLower(hash)) {
            fmt.Printf(color.Cyan("The password is: ") + color.Green(line) + "\n")

        }
    }
    fmt.Printf(color.Cyan("Cracking for ") + color.Yellow(hash) + color.Cyan(" ended\n"))
}
