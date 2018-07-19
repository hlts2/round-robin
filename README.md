# round-robin
round-robin is balancing algorithm written in golang

## Requrement
Go (>= 1.8)

## Installation

```shell
go get github.com/hlts2/round-robin
```

## Example

```go
rr, err := roundrobin.New([]string{
    "server-1",
    "server-2",
    "server-3",
})

rr.Next() // server-1
rr.Next() // server-2
rr.Next() // server-3
rr.Next() // server-1
rr.Next() // server-2
rr.Next() // server-3
```
## Author
[hlts2](https://github.com/hlts2)

## LICENSE
round-robin released under MIT license, refer [LICENSE](https://github.com/hlts2/round-robin/blob/master/LICENSE) file.
