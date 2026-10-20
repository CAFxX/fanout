package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 1
    var v3 uint16 = 65535
    var v4 bool = (127 >= 42)
    v5 := make(map[int]bool)
    switch "" {
    case "\n\t":
        v6 := GenS1{F0: 2147483647}
        _ = v6
    default:
        v7 := GenS2{F0: v2, F1: 42, F2: 9223372036854775807}
        _ = v7
    }
    delete(v5, 128)
    var v8 uint8 = 111
    v8 = (v8 - (v8 - 42))
    v9 := GenS1{F0: 42}
    v2 = v2
    v10 := GenS2{F0: (v2 & v2), F1: 100, F2: -767}
    v8 = v8
    var v11 int = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
