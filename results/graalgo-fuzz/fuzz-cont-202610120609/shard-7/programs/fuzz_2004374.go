package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint32 = 0
    v2 := []int{0, 174, 128}
    var v3 float32 = 1.5
    if (true && false) {
        _ = len(v2)
    } else {
        fmt.Println(v3, v1)
        v4 := GenS1{F0: 42, F1: 496, F2: 921}
        _ = v4
    }
    var v5 float32 = 1.5
    var v6 int = 9223372036854775807
    v6 += -2147483648
    v5 += ((v3 * v3) + v3)
    fmt.Println(v6, v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        v7 := GenS1{F0: 2, F1: -1, F2: 2}
        v8 := GenS2{F0: 4294967295}
        _ = v7
        _ = v8
    } else {
        if (!false) {
            fmt.Println(v3, v6)
            var v9 uint32 = 427
            _ = v9
        } else {
            v1 -= (v1 << ((0) & 63))
            var v10 float64 = 1e10
            _ = v10
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
