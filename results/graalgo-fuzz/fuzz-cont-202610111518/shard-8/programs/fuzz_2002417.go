package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int64 }

func main() {
    var v1 float32 = float32(((float64(((0.1 * -2.25) + 0.1)) * (-2.25 + ((1.5 + 1e10) - 3.14159))) + 0.1))
    var v2 uint64 = 24
    var v3 string = ""
    var v4 int8 = 100
    var v5 uint16 = 596
    if true {
        v6 := []int{2147483647, -772}
        var v7 int8 = v4
        _ = v6
        _ = v7
    }
    fmt.Println(v3)
    if (!false) {
        v8 := []int{516, -604}
        _ = v8
    }
    switch v3 {
    case "hello":
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint = 1
            v11 := []int{}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 uint = 4294967295
        if (0.0 >= (3.14159 + -2.25)) {
            fmt.Println(v13, fi12, v2)
        } else {
            if (v4 < (v4 >> (((^(v13))) & 63))) {
                v1 -= v1
            }
        }
        var v14 uint64 = (v2 * (^(v2)))
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(fi15, v3, v1)
        v5 -= 498
        _ = fi15
    }
    var v16 float32 = 0.1
    v17 := GenS1{F0: uint8(v5), F1: v5, F2: int64(v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
