package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint16 = 436
    var v2 int8 = -57
    var v3 int32 = 2147483647
    var v4 uint16 = v1
    if (100 <= 100) {
        v5 := []int{2147483647, 9223372036854775807, 32767, 0}
        _ = v5
    }
    if (!(true && false)) {
        v6 := []int{127, 2147483647, -27, 4294967295}
        var v7 uint8 = 189
        _ = v6
        _ = v7
    }
    if true {
        if (2 <= 42) {
            var v8 uint16 = (v4 ^ 255)
            _ = v8
        }
    }
    var v9 float32 = 1.5
    v4 += v4
    v1 -= (v1 / (((^(v4))) | 1))
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: (v3 | (^(v3)))}
        fmt.Println(fi10, v3)
        _ = fi10
        _ = v11
    }
    v9 += 0.0
    v9 += v9
    v12 := []int{601, 32768, 42}
    if (2 <= 0) {
        fmt.Println(v2, v2, v2)
    }
    var v13 int8 = (v2 & v2)
    v9 += (v9 - v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
