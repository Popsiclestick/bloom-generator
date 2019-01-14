package main

import (
	"bufio"
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/willf/bloom"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var file string
	flag.StringVar(&file, "f", "bloom-input", "New line delimited file")

	var bits uint
	flag.UintVar(&bits, "b", 1, "Bits (m)")

	var hashes uint
	flag.UintVar(&hashes, "h", 1, "Hashing functions (k)")

	flag.Parse()

	log.Info("Initializing & loading bloom filter")
	filter := bloom.New(bits, hashes)

	inputFile, err := os.Open(file)
	check(err)
	defer inputFile.Close()

	// Scan the file by line and load the bloom filter
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		filter.Add(scanner.Bytes())
	}

	bloomfilterFile, err := os.Create("bloomfilter")
	check(err)
	defer bloomfilterFile.Close()

	filter.WriteTo(bloomfilterFile)
	log.Info("Bloomfilter written to file")
}
