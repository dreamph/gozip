# gozip

Golang Zip File
- Simple, Easy
- Support File or Directory

Install
=======
```go
package main

import (
	"github.com/dreamph/gozip"
	"log"
)

func main() {
	err := gozip.Zip("example/output/test.zip", []string{"example/dir1", "example/test-file.txt"})
	if err != nil {
		log.Fatal("Zip file error", err)
	}

	err = gozip.Zip("example/output/test-with-password.zip", []string{"example/dir1", "example/test-file.txt"}, "password")
	if err != nil {
		log.Fatal("Zip file with password error", err)
	}
}
```

Buy Me a Coffee
=======
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dreamph)