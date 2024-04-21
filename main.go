package main 

import (
	"flag"
	"fmt"

	hex "./hex"
	binary "./binary"

	fileutils "./fileutils"
)

const (
	hexdump_creation_date = "2024"
	hexdump_version = "0.2"

	default_silent_mode = false
	default_bytes_per_row = 8

	// data dump modes
	dump_as_hex = 0
	dump_as_binary = 1
)

var (
	silentMode bool = default_silent_mode
	bytesPerRow int = default_bytes_per_row
	dumpMode int = dump_as_hex
)

func showBanner() {
	fmt.Printf ("FileDumper Utility %s\n", hexdump_version)
	fmt.Printf ("(c) Colin Wilcox %s\n\n", hexdump_creation_date)
}

func parseArguments() {

	var binaryMode bool

	flag.BoolVar(&silentMode, "silent", default_silent_mode, "Utility operating mode.")
	flag.IntVar(&bytesPerRow, "width", default_bytes_per_row, "Number of bytes to display per row.")
	flag.BoolVar(&binaryMode, "binary", false, "Dump file contents in binary.")
	flag.Parse()

	// dump as specified with default as hex
	if binaryMode {
		dumpMode = dump_as_binary
	} else {
		dumpMode = dump_as_hex
	}

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
		switch(dumpMode) {
			case dump_as_hex: hex.Dump(data, bytesPerRow)
			break;
			case dump_as_binary: binary.Dump(data, bytesPerRow)
			break;
		}
	} else {
		fmt.Println(err)
	}
}