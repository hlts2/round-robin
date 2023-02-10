# round-robin

round-robin is balancing algorithm written in golang (with generics)

## Installation

```shell
go get github.com/skamenetskiy/round-robin
```

## Example

```go
package main

import (
	"net/url"

	roundrobin "github.com/skamenetskiy/round-robin"
)

func main() {
	rr, _ := roundrobin.New(
		&url.URL{Host: "192.168.33.10"},
		&url.URL{Host: "192.168.33.11"},
		&url.URL{Host: "192.168.33.12"},
		&url.URL{Host: "192.168.33.13"},
	)

	rr.Next() // {Host: "192.168.33.10"}
	rr.Next() // {Host: "192.168.33.11"}
	rr.Next() // {Host: "192.168.33.12"}
	rr.Next() // {Host: "192.168.33.13"}
	rr.Next() // {Host: "192.168.33.10"}
	rr.Next() // {Host: "192.168.33.11"}	
}
```

## Author

[hlts2](https://github.com/hlts2)

## LICENSE

round-robin released under MIT license, refer [LICENSE](https://github.com/hlts2/round-robin/blob/master/LICENSE) file.
