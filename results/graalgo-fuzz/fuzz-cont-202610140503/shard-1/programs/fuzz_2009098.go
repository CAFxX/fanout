package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int8 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]bool)
    var v3 int8 = 42
    var v4 int64 = -9223372036854775808
    fmt.Println(v3, v1, v4)
    var v5 bool = (221 != 2)
    var v6 uint = 4294967295
    fmt.Println(v4)
    v7 := GenS1{F0: 221, F1: (v6 << ((v6) & 63)), F2: (^(v3))}
    var v8 uint16 = uint16(v6)
    if (!v5) {
        v4 = -877
    } else {
        switch "foo bar" {
        case "foo bar":
            v4 -= (-(v4))
        case "":
            delete(v2, 2)
        }
    }
    v9 := GenS1{F0: 137, F1: v6, F2: 100}
    v3 -= (-(v3))
    v5 = (!v5)
    var v10 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
