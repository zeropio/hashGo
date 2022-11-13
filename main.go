package main

import (
	"bufio"
	"flag"
	"fmt"
	h "hashgo/hashes"
	"os"

	"github.com/labstack/gommon/color"
)

func main() {

    // Flags
    hashfile := flag.String("f", "", "Select the file with the hash or hashes")
    wordlist := flag.String("w", "", "Select the wordlist it will use")
    hash := flag.String("h", "", "Select your hash type")
    flag.Parse()

    // File error handling
    hFile, err := os.Open(*hashfile)
    if err != nil{
        text := color.Red("Error opening " + *hashfile + "\nFile missing?" + "\n")
        fmt.Printf(text)
    }
    wFile, err := os.Open(*wordlist)
    if err != nil{
        text := color.Red("Error opening " + *wordlist + "\nFile missing?" + "\n")
        fmt.Printf(text)
    }

    // Open hash file and send with the wordlist
    readFile := bufio.NewScanner(hFile)
        readFile.Split(bufio.ScanLines)
        for readFile.Scan() {
            line := readFile.Text()
            switch {
            case string(*hash) == "md5":
                h.Md5(line, wFile)
            case string(*hash) == "md4":
                h.Md4(line, wFile)
            case string(*hash) == "sha1":
                h.Sh1(line, wFile)
            case string(*hash) == "ntlm":
                h.Ntlm(line, wFile)
            default:
                fmt.Printf("Please select a hash tpye\n")
            }
            
        }
    }

