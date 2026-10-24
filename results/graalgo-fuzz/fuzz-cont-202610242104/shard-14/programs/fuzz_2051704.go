package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 int = 42
    var v2 uint = 65535
    v3 := []uint32{0}
    var v4 int32 = 2147483647
    var v5 int8 = int8(v2)
    var v6 uint16 = 0
    if false {
        _ = len(v3)
    }
    v7 := GenS1{F0: 0, F1: 2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch ("" + ("\n\t" + "hello")) {
    case "foo bar":
        var v8 int64 = -142
        _ = v8
    default:
        var v9 int8 = v5
        _ = v9
    }
    v6 -= v6
    v6 -= (^(v6))
    v1 += v1
    v10 := GenS1{F0: 566, F1: uint8(v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
