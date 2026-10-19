package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }

func main() {
    var v1 int16 = -1
    var v2 uint64 = 4294967295
    var v3 int = 42
    v4 := make(map[int]string)
    switch -9223372036854775808 {
    case -1:
        v1 = int16(v3)
    case -979:
        var v5 uint16 = 65535
        _ = v5
    }
    var v6 float32 = 0.0
    _ = len(v4)
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v6, v6)
    case "hello":
        var v7 uint64 = v2
        _ = v7
    default:
        v3 -= (^(v3))
    }
    v8 := GenS1{F0: 4294967295, F1: uint32(v2), F2: 100}
    v1 = (^(v1))
    var v9 string = ("foo bar" + "\n\t")
    var v10 string = v9
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
