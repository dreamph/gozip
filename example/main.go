package main

import (
	"github.com/dreamph/gozip"
	"log"
)

func main() {
	err := gozip.Zip("example/output/test.zip", []string{"example/dir1", "example/test-file.txt"})
	if err != nil {
		log.Fatal("Zip file error.", err)
	}

	err = gozip.Zip("example/output/test-with-password.zip", []string{"example/dir1", "example/test-file.txt"}, "password")
	if err != nil {
		log.Fatal("Zip file error.", err)
	}
}
