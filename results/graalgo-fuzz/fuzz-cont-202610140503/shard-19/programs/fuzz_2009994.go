package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 uint; F1 uint; F2 int32 }

func main() {
    var v1 uint8 = 51
    var v2 int8 = 66
    v3 := []uint32{}
    v4 := []uint32{2147483648, 0, 749, 2147483648}
    var v5 uint64 = 1
    v6 := []uint32{}
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v7 := GenS1{F0: (^(v2)), F1: 0, F2: ""}
    var v8 int = 681
    switch "foo bar" {
    case "\n\t":
        var v9 int16 = 1000
        _ = v9
    case "foo bar":
        fmt.Println(v2)
    }
    v10 := GenS1{F0: v2, F1: 4294967295, F2: ""}
    if ((!((-188 > 0) && true)) || ((v1 < (v1 % ((2) | 1))) || false)) {
        var v11 uint8 = 0
        _ = v11
    } else {
        v2 = v2
    }
    v8 = int(v2)
    fmt.Println(v5)
    if true {
        v12 := GenS2{F0: 943, F1: uint(v1), F2: -1}
        v2 -= (v2 - v2)
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
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
}
