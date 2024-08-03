# gozip

Golang Zip File
- Simple, Easy
- Support File or Directory

Example
=======
```go
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
```

Buy Me a Coffee
=======
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dreamph)