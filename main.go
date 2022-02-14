package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var out io.Writer = os.Stdout

func main() {

	cmp := flag.NewFlagSet("cmp", flag.ExitOnError)

	first := cmp.String("first", "", "name of the first file to be compared")
	second := cmp.String("second", "", "name of the second file file to be compared")

	args := os.Args[1:]
	cmp.Parse(args)

	if len(args) != 4 {
		cmp.PrintDefaults()
	} else {

		if _, err := checkIfValidFiles(*first, *second); err != nil {
			exitGracefully(err)
		}

		eq, err := equals(*first, *second)
		if err != nil {
			exitGracefully(err)
		}

		if eq {
			fmt.Fprintln(out, "true")
		} else {
			fmt.Fprintln(out, "false")
		}

	}
}

func checkIfValidFiles(filepath1, filepath2 string) (bool, error) {
	if _, err := checkIfValidFile(filepath1); err != nil {
		return false, err
	}

	if _, err := checkIfValidFile(filepath2); err != nil {
		return false, err
	}
	return true, nil
}

func checkIfValidFile(filename string) (bool, error) {
	isJsonFile := filepath.Ext(filename) == ".json"
	if !isJsonFile {
		return false, fmt.Errorf("file %s is not json", filename)
	}

	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exist", filename)
	}

	return true, nil
}

func getHashFromFile(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}

	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func equals(filepath1, filepath2 string) (bool, error) {
	var err error
	hash1, err := getHashFromFile(filepath1)
	if err != nil {
		return false, err
	}
	hash2, err := getHashFromFile(filepath2)
	if err != nil {
		return false, err
	}

	return (hash1 == hash2), nil
}

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
