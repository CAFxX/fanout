package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool; F1 int8 }
func genFn0() uint8 {
    v5 := []int{1, 0, 109}
    var v6 uint16 = 255
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint32 = 365
    v2 := []int{4294967295, 0}
    var v3 uint8 = 42
    var v4 int8 = -128
    v7 := GenS1{F0: -2147483648, F1: 280}
    v4 -= (v4 / (((^(v4))) | 1))
    var v8 int64 = 333
    if false {
        v9 := GenS2{F0: false, F1: (v4 ^ (^(v4)))}
        _ = v9
    }
    if (!((v3 * 2) <= v3)) {
        var v10 int32 = 42
        _ = v10
    }
    if (!false) {
        v11 := GenS1{F0: -805, F1: 1}
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v4 -= (v4 % (((^(v4))) | 1))
        _ = fi12
    }
    v1 = (v1 >> ((255) & 63))
    var v13 uint32 = (v1 & (v1 >> ((55) & 63)))
    v1 += (v1 + (v13 % ((v13) | 1)))
    var v14 int32 = 42
    fmt.Println(v1, v13)
    _ = len(v2)
    v15 := GenS2{F0: false, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
