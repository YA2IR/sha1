# SHA1 Implementation in Go
This is an implementation of [SHA-1](https://en.wikipedia.org/wiki/SHA-1) in Go. It follows the [FIPS 180-1](https://nvlpubs.nist.gov/nistpubs/Legacy/FIPS/fipspub180-1.pdf) standard.

## Features
-  Interactive REPL for experimentation (see [As a REPL](#as-a-repl))
## Limitations
- This implementation is strictly for educational purposes. SHA-1 is considered cryptographically broken and should **not** be used in production.
## Usage

### As a Library
```go
package main

import (
	"fmt"
	"github.com/YA2IR/sha1"
)

func main() {
	s := sha1.NewSHA1()

	message1 := "hello world"
	message2 := "hello"

	digest := s.Hash([]byte(message1))
	fmt.Printf("SHA1 of \"%s\" = %x\n", message1, digest)

	digest = s.Hash([]byte(message2))
	fmt.Printf("SHA1 of \"%s\" = %x\n", message2, digest)

}
```
output:
```
SHA1 of "hello world" = 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
SHA1 of "hello" = aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
```
### As a REPL
Clone this repo:
```bash
git clone https://github.com/YA2IR/sha1.git
cd sha1
```
Then you can run:
```bash
go run cmd/sha1/main.go
```
```
Welcome to SHA1 REPL, write 'exit' to exit
> test
SHA1: a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
> test2
SHA1: 109f4b3c50d7b0df729d299bc6f8e9ef9066971f
> exit

```

