package main

import (
	"fmt"
	"os"

	"github.com/codingsince1985/checksum"
)

func main() {

	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	arg1 := os.Args[1]

	switch arg1 {
	case "-p", "--print":
		if len(os.Args) == 4 {
			fmt.Println(fileSum(os.Args[2], os.Args[3]))
		} else {
			help()
		}
	case "-c", "--compare":
		if len(os.Args) == 5 {
			if fileSum(os.Args[2], os.Args[3]) == os.Args[4] {
				fmt.Println("Valid. File hash validated.")
			} else {
				fmt.Println("Invalid: File might have been corrupted.")
			}
		} else {
			help()
		}
	case "-h", "--help":
		help()
	default:
		help()
	}
}

func fileSum(file string, sumType string) string {
	switch sumType {
	case "md5", "MD5":
		s, _ := checksum.MD5sum(file)
		return s
	case "sha256", "SHA256":
		s, _ := checksum.SHA256sum(file)
		return s
	case "sha1", "SHA1":
		s, _ := checksum.SHA1sum(file)
		return s
	case "crc32", "CRC32":
		s, _ := checksum.CRC32(file)
		return s
	case "blake2s256":
		s, _ := checksum.Blake2s256(file)
		return s
	}
	return "invalid parameter for checksum"
}

func help() {
	fmt.Println("sum - checksum a file or compare a given hash")
	fmt.Println(" ")
	fmt.Println("sum OPTION FILE [HASHTYPE] [HASH]")
	fmt.Println(" ")
	fmt.Println("DESCRIPTION")
	fmt.Println(" ")
	fmt.Println("-p, --print FILE HASHTYPE")
	fmt.Println("     create checksum from a FILE and prints the hash to the console.")
	fmt.Println(" ")
	fmt.Println("-c, --compare FILE HASHTYPE HASH")
	fmt.Println("     compare the given hash with the checksum from FILE")
	fmt.Println("     and prints to the console, if it's VALID or NOT VALID.")
	fmt.Println("     (NOT valid means file have been corrupted)")
	fmt.Println(" ")
	fmt.Println("-h, --help display this help and exit ")
	fmt.Println(" ")
}
