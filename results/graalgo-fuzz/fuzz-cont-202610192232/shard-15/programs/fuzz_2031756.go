package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 uint64 = 42
    var v2 uint16 = 32767
    v3 := []string{"foo bar"}
    switch "" {
    case "hello":
        fmt.Println(v2, v2)
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "\n\t":
        v4 := GenS1{F0: 311, F1: true, F2: uint8(v2)}
        _ = v4
    }
    var v5 uint8 = 1
    v5 += (^(v5))
    v5 = (v5 | v5)
    v2 += (v2 << ((2) & 63))
    v6 := GenS1{F0: -253, F1: (-1 == -32768), F2: uint8(v2)}
    v7 := GenS1{F0: -1, F1: true, F2: (v5 * v5)}
    var v8 string = "\n\t"
    v8 = ("hello" + v8)
    var v9 int = 42
    v10 := GenS2{F0: v8, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
