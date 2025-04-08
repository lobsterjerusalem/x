package main
import (
	"os"
	"strings"
	"encoding/hex"
	"fmt"
)

func main(){
	var hexDataIn string
	var bin []byte
	filePrefix := "output"
	if len(os.Args) < 3 {
		fmt.Println(fmt.Sprintf("Usage: %s <hex or hex filepath> <optional: output filepath prefix>", os.Args[0]))
		os.Exit(1)
	}
	hexDataIn = os.Args[1]
	filePrefix = os.Args[2]

	// check if a filepath was provided
	_,err := os.Stat(hexDataIn)
	if err == nil {
		// file exists
		hexDataBytes, err := os.ReadFile(hexDataIn)
		if err != nil {
			fmt.Println(fmt.Sprintf("error reading file: %s", err))
			os.Exit(1)
		}
		hexDataIn = strings.TrimSpace(string(hexDataBytes))
		bin,_ = hex.DecodeString(hexDataIn)
		if err != nil {
			fmt.Println(fmt.Sprintf("error decoding string arg: %s", err))
			os.Exit(1)
		}
	} else { // probably a string
		fmt.Println(fmt.Sprintf("err opening %s, %s", hexDataIn, err))
		bin,err = hex.DecodeString(strings.TrimSpace(hexDataIn))
		if err != nil {
			fmt.Println(fmt.Sprintf("error decoding string arg: %s", err))
			os.Exit(1)
		}
	}

	fileNameBin := fmt.Sprintf("%s.bin", filePrefix)
	fileNameHexStr := fmt.Sprintf("%s.hexstr", filePrefix)

	os.WriteFile(fileNameBin, bin, 0600)
	fmt.Println(fmt.Sprintf("saved as binary to %s", fileNameBin))


	hexStr := makeHexStr(hexDataIn)
	os.WriteFile(fileNameHexStr, []byte(hexStr), 0600)
	fmt.Println(fmt.Sprintf("saved as hex str to %s", fileNameHexStr))
}

func makeHexStr(hexData string) string {
	var result strings.Builder
	result.WriteString("\"")
	for i := 0; i< len(hexData); i+= 2 {
		hexByte := hexData[i:i+2]
		result.WriteString(fmt.Sprintf("\\x%s", hexByte))
		if i % 32 == 0 && i != 0 {
			result.WriteString("\"+\n\"")
		}
	}
	result.WriteString("\"\n")

	return strings.ToLower(result.String())
}
