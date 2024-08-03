package main

import (
	"github.com/dreamph/gozip"
	"log"
)

func main() {
	// Zip
	err := gozip.Zip("example/output/test.zip", []string{"example/dir1", "example/test-file.txt"})
	if err != nil {
		log.Fatal("Zip file error.", err)
	}

	// Zip with password
	err = gozip.Zip("example/output/test-with-password.zip", []string{"example/dir1", "example/test-file.txt"}, "password")
	if err != nil {
		log.Fatal("Zip file error.", err)
	}

	// Unzip
	err = gozip.Unzip("example/output/test.zip", "example/output/test")
	if err != nil {
		log.Fatal("Zip file error.", err)
	}

	// Unzip with password
	err = gozip.Unzip("example/output/test-with-password.zip", "example/output/test-with-password", "password")
	if err != nil {
		log.Fatal("Zip file error.", err)
	}
}
