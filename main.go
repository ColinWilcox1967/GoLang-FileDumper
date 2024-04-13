package main 

import (
	"flag"
	"fmt"

	hex "./hex"
	fileutils "./fileutils"
)

const (
	hexdump_creation_date = "2024"
	hexdump_version = "0.1"

	default_silent_mode = false
	default_bytes_per_row = 8
)

var (
	silentMode bool = default_silent_mode
	bytesPerRow int = default_bytes_per_row
)

func showBanner() {
	fmt.Printf ("HexDump Utility %s\n", hexdump_version)
	fmt.Printf ("(c) Colin Wilcox %s\n\n", hexdump_creation_date)
}

func parseArguments() {
	flag.BoolVar(&silentMode, "silent", default_silent_mode, "Utility operating mode.")
	flag.IntVar(&bytesPerRow, "width", default_bytes_per_row, "Number of bytes to display per row.")
	flag.Parse()

	if bytesPerRow <= 0 {
		bytesPerRow = default_bytes_per_row
	}
}

func main() {
	parseArguments()

	if !silentMode {
		showBanner()
	}

	data, err := fileutils.ReadFile("test.txt")
	if err == nil {
		hex.DumpAsHex(data, bytesPerRow)
	} else {
		fmt.Println(err)
	}
}