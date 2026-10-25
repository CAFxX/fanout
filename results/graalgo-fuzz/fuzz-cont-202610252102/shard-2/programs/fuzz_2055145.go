package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }

func main() {
    var v1 int8 = 109
    var v2 int8 = (v1 << ((18446744073709551615) & 63))
    v3 := make(map[string]uint32)
    var v4 int16 = 1
    var v5 uint = 4294967295
    var v6 uint = v5
    v7 := []int{}
    v8 := v3["a"]
    fmt.Println(v6)
    var v9 uint16 = 2
    if ((v5 % ((179) | 1)) < (v6 & (v5 >> ((v6) & 63)))) {
        v10 := GenS1{F0: 0, F1: (v2 % ((v2) | 1)), F2: 0}
        v5 += 1
        v1 -= (v2 ^ v1)
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: -9223372036854775808, F1: v1, F2: -461}
    v5 = uint(v8)
    fmt.Println(v9, v4)
    if true {
        var v14 string = "\n\t"
        var v15 uint16 = 0
        var v16 int = 32768
        _ = v14
        _ = v15
        _ = v16
    }
    v1 = (v1 / ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
