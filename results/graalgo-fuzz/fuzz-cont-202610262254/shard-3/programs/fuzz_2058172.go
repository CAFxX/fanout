package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 18446744073709551615
    var v2 bool = (!false)
    var v3 int = -670
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: 898, F1: 1, F2: fi4}
        if (v1 == v1) {
            fmt.Println(v1, v1)
            v6 := []int{-9223372036854775808, v3, fi4}
            _ = v6
        } else {
            var v7 int32 = 1
            switch ("a" + "foo bar") {
            case "":
                var v8 float64 = (1.5 - 0.1)
                _ = v8
            case "\n\t":
                v9 := GenS2{F0: (v1 >> (((v1 + v1)) & 63))}
                _ = v9
            case "a":
                fmt.Println(v2, v2, v2)
            }
            v7 -= (v7 / ((v7) | 1))
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v10 := []int{v3}
    var v11 int64 = 0
    var v12 uint = (v1 / (((v1 % ((18446744073709551615) | 1))) | 1))
    for fi13 := 0; fi13 < 11; fi13++ {
        v11 -= (^(v11))
        v12 = v12
        var v14 int16 = 42
        _ = fi13
        _ = v14
    }
    var v15 int16 = -908
    var v16 uint64 = 131
    v1 = 255
    for fi17 := 0; fi17 < 4; fi17++ {
        fmt.Println(v16, v15)
        _ = fi17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        fmt.Println(v12, v16, v1)
        _ = fi18
    }
    v19 := GenS2{F0: (v12 << ((v12) & 63))}
    fmt.Println(v3, v1)
    _ = len(v10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
