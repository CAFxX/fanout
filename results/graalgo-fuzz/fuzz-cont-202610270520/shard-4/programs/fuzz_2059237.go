package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    var v1 uint = 42
    var v2 int = -2147483648
    var v3 int16 = -25
    var v4 uint8 = 42
    var v5 uint8 = v4
    var v6 uint16 = 255
    v5 = (^(v4))
    var v7 float32 = 0.0
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint = 4294967295
        v10 := GenS1{F0: 100, F1: int64(v5)}
        v2 += (fi8 ^ -325)
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 bool = (!false)
    v12 := GenS1{F0: int8(v3), F1: 9223372036854775807}
    fmt.Println(v7)
    switch (v2 >> (((^(v1))) & 63)) {
    case 128:
        v4 = 255
    case 1:
        v2 -= v2
    case -2147483648:
        v4 = (v4 ^ (^(v5)))
    default:
        v7 = 1e10
    }
    v3 = v3
    v1 = (v1 ^ (^(v1)))
    var v13 uint8 = v4
    v14 := GenS2{F0: "a", F1: (^(v3))}
    if true {
        v15 := []int{v2}
        _ = v15
    } else {
        v13 = v5
    }
    v16 := []int{(v2 | 77), v2}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
