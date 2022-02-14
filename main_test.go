package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCLI(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		output string
	}{
		{"equal files", []string{"./cmp", "-first", "./test_files/file1.json", "-second", "./test_files/file1.json"}, "true\n"},
		{"different files", []string{"./cmp", "-first", "./test_files/file1.json", "-second", "./test_files/file2.json"}, "false\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			out = bytes.NewBuffer(nil)
			main()
			if actual := out.(*bytes.Buffer).String(); actual != tt.output {
				fmt.Println(actual, tt.output)
				t.Errorf("expected %s, but got %s", tt.output, actual)
			}
		})
	}
}

func TestEqual(t *testing.T) {

	tests := []struct {
		name    string
		file1   string
		file2   string
		want    bool
		wantErr bool
	}{
		{"equal files", "./test_files/file1.json", "./test_files/file1.json", true, false},
		{"different files", "./test_files/file1.json", "./test_files/file2.json", false, false},
		{"non json file", "./test_files/file1.txt", "./test_files/file2.json", false, true},
		{"non existent file", "file1.txt", "./test_files/file2.json", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := equals(tt.file1, tt.file2)

			if (err != nil) != tt.wantErr {
				t.Errorf("equals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckIfValidFiles(t *testing.T) {
	file1, err := os.Open("./test_files/file1.json")
	if err != nil {
		panic(err)
	}

	defer file1.Close()

	file2, err := os.Open("./test_files/file2.json")
	if err != nil {
		panic(err)
	}

	defer file2.Close()

	tests := []struct {
		name      string
		filename1 string
		filename2 string
		want      bool
		wantErr   bool
	}{
		{"Both files do exist", file1.Name(), file2.Name(), true, false},
		{"Both files do not exist", "nowhere/file1.json", "nowhere/file2.json", false, true},
		{"File does not exist", file1.Name(), "nowhere/test.json", false, true},
		{"File is not csv", file1.Name(), "test.txt", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkIfValidFiles(tt.filename1, tt.filename2)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkIfValidFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkIfValidFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckIfValidFile(t *testing.T) {
	file, err := ioutil.TempFile("", "test*.json")
	if err != nil {
		panic(err)
	}

	defer os.Remove(file.Name())

	tests := []struct {
		name     string
		filename string
		want     bool
		wantErr  bool
	}{
		{"File does exist", file.Name(), true, false},
		{"File does not exist", "nowhere/test.json", false, true},
		{"File is not csv", "test.txt", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkIfValidFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkIfValidFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkIfValidFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
