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
next, err := RoundRobin(Servers {
    "server-1",
    "server-2",
    "server-3",
})

next() // server-1
next() // server-2
next() // server-3
next() // server-1
next() // server-2
next() // server-3
```
