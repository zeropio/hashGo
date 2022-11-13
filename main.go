package main

import (
	"flag"
	"fmt"
	h "hashgo/hashes"
	"os"

	"github.com/labstack/gommon/color"
)

func flagHelp() {
    fmt.Printf(color.Cyan("Files\n"))
    fmt.Printf(`-f file
    Path to the file with hashes`)
    fmt.Printf("\n")
    fmt.Printf(`-w wordlist
    Path to the wordlist`)
    fmt.Printf(color.Cyan("\nHashes\n"))
    fmt.Printf(`-t type
    Select the type of hash (md5, sha1, ntlm,...). Lowercase`)
    fmt.Printf(color.Cyan("\nOther\n"))
    fmt.Printf(`-h
    Print help`)
    fmt.Printf(color.Cyan("\nBy Zeropio\n"))
}

func main() {
    // var
    var help bool

    // Flags
    input := flag.String("f", "", "Input the hash")
    wordlist := flag.String("w", "", "Select the wordlist it will use")
    hType := flag.String("t", "", "Select your hash type")
    flag.BoolVar(&help, "h", false, "Print help")
    flag.Parse()

    // Help
    if (help) {
        flagHelp()
        os.Exit(0)
    }

    // File error handling
    wFile, err := os.Open(*wordlist)
    if err != nil{
        fmt.Printf(color.Red("Error opening wordlist: " + *wordlist + "\nFile missing?\n"))
        flagHelp()
        os.Exit(0)
    }

    // Select hash and send with the wordlist
    switch {
    case string(*hType) == "md5":
        h.Md5(*input, wFile)
    case string(*hType) == "md4":
        h.Md4(*input, wFile)
    case string(*hType) == "sha1":
        h.Sh1(*input, wFile)
    case string(*hType) == "ntlm":
        h.Ntlm(*input, wFile)
    default:
        fmt.Printf(color.Red("Please select a valid hash tpye\n"))
        flagHelp()
    }        
}


