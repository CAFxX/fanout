package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint8 = 220
    var v2 uint64 = 0
    var v3 int16 = 42
    var v4 int = 127
    if true {
        switch "\n\t" {
        case "a":
            v4 = v4
        case "foo bar":
            v5 := GenS1{F0: 742, F1: 54}
            _ = v5
        case "":
            fmt.Println(v3, v1, v3)
        }
    } else {
        v2 += (v2 << ((1) & 63))
        v6 := []int{(v4 & 2147483647), (v4 + -362), (v4 ^ v4)}
        _ = v6
    }
    v4 = v4
    for fi7 := 0; fi7 < 12; fi7++ {
        v2 = (v2 ^ 18446744073709551615)
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := []int{v4, fi7, fi7}
            fmt.Println(v3)
            v1 = (v1 << ((833) & 63))
            v1 -= v1
            _ = fi8
            _ = v9
        }
        if (!true) {
            v10 := GenS2{F0: -999}
            _ = v10
        }
        _ = fi7
    }
    fmt.Println(v4, v3)
    v2 -= 0
    v1 = v1
    var v11 int64 = -9223372036854775808
    v11 = (v11 | v11)
    v12 := []int{}
    if false {
        v2 = (v2 ^ v2)
        v13 := GenS2{F0: -52}
        var v14 int8 = -1
        _ = v13
        _ = v14
    }
    v1 -= 0
    if true {
        var v15 uint16 = 42
        fmt.Println(v3, v3)
        _ = v15
    }
    v2 += 4294967295
    if false {
        v2 = (v2 * 18446744073709551615)
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
