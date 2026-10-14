package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }

func main() {
    var v1 int32 = 752
    var v2 int8 = 100
    v3 := []uint32{2, 940, 0}
    var v4 uint8 = 108
    switch 4294967295 {
    case -9223372036854775808:
        var v5 float32 = -2.25
        _ = v5
    case 2147483647:
        var v6 bool = true
        _ = v6
    }
    fmt.Println(v2, v2)
    var v7 float64 = (0.0 + 3.14159)
    var v8 uint32 = 905
    for fi9 := 0; fi9 < 6; fi9++ {
        switch 189 {
        case 2147483647:
            var v10 uint = 0
            _ = v10
        }
        var v11 float64 = -2.25
        _ = fi9
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        if (v4 >= (v4 << ((18446744073709551615) & 63))) {
            v13 := GenS1{F0: "hello", F1: 1}
            var v14 int32 = 0
            _ = v13
            _ = v14
        } else {
            fmt.Println(v7, v4)
            v8 -= v8
            for fi15 := 0; fi15 < 1; fi15++ {
                v4 = (v4 % ((v4) | 1))
                v1 = (v1 + (v1 * (v1 | (v1 >> ((2) & 63)))))
                _ = fi15
            }
        }
        _ = fi12
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi16
    }
    v3 = append(v3, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
