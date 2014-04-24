bencode
=======

Bencode implementation in Go

## Download

```bash
$ go get github.com/marksamman/bencode
```

## Usage

### Encode
bencode.Encode takes a map[string]interface{} as argument and returns a byte array. Example:
```go
package main

import (
	"fmt"

	"github.com/marksamman/bencode"
)

func main() {
	dict := make(map[string]interface{})
	dict["string key"] = "hello world"
	dict["int key"] = 123456
	fmt.Printf("bencode encoded dict: %s\n", bencode.Encode(dict))
}
```

If a specific order is desired for dictionary keys when generating a bencode string, specify the keys in desired order in the dictionary key "\_\_keys". For dictionaries generated from bencode.Decode, "\_\_keys" is already set to the order the keys were read from the file. Example:
```go
dict["__keys"] = []string{"int key", "string key"}
```

### Decode
bencode.Decode takes an io.Reader as argument and returns (map[string]interface{}, error). Example:
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marksamman/bencode"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dict, err := bencode.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("string: %s\n", dict["string_key"].(string))
	fmt.Printf("int: %d\n", dict["int_key"].(int64))
}
```
