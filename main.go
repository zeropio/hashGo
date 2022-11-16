package main

import (
	"flag"
	"fmt"
	h "hashgo/hashes"
	"os"

	"github.com/labstack/gommon/color"
)

func flagHelp() {
    fmt.Printf(color.Cyan("\nHash\n"))
    fmt.Printf(`-f file
    Path to the file with hashes`)
    fmt.Printf("\n")
    fmt.Printf(`-w wordlist
    Path to the wordlist`)
    fmt.Printf("\n")

    fmt.Printf(color.Cyan("\nHashes\n"))
    fmt.Printf(`-t type
    Select the type of hash (md5, sha1, ntlm,...). Lowercase`)
    fmt.Printf(color.Cyan("\nOther\n"))
    fmt.Printf(`-v
    Print verbose. Print tested passwords`)
    fmt.Printf("\n")
    fmt.Printf(`-h
    Print help`)
    fmt.Printf("\n")

    fmt.Printf(color.Cyan("\nBy Zeropio\n") + color.Red("with Love"))
}

func main() {
    // var
    var help bool

    // Flags
    input := flag.String("f", "", "Input the hash")
    wordlist := flag.String("w", "", "Select the wordlist it will use")
    verbose := flag.Bool("v", false, "Set verbose (print tested password)")
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
        h.ToMD5(*input, wFile, *verbose)
    case string(*hType) == "md4":
        h.ToMD4(*input, wFile, *verbose)
    case string(*hType) == "sha1":
        h.ToSH1(*input, wFile, *verbose)
    case string(*hType) == "ntlm":
        h.ToNTLM(*input, wFile, *verbose)
    default:
        fmt.Printf(color.Red("Please select a valid hash tpye\n"))
        flagHelp()
    }        
}


