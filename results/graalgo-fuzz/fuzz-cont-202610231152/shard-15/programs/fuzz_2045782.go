package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{2147483647, 32767}
    var v3 float64 = ((0.1 - (((-2.25 - 0.1) + 1e10) - float64((0.0 + ((-2.25 * float32(-2.25)) * (1e10 + 0.1)))))) - float64(1e10))
    var v4 float32 = -2.25
    v5 := make(map[int]bool)
    var v6 uint8 = 255
    v7 := GenS1{F0: (v6 | 28), F1: 42}
    if false {
        var v8 uint = 4294967295
        v8 = (v8 >> (((v8 + (v8 / (((v8 * 464)) | 1)))) & 63))
        _ = v8
    } else {
        v1 += (v1 + (^(v1)))
    }
    if (((v1 << ((1) & 63)) == (v1 % ((0) | 1))) && true) {
        v6 -= v6
        var v9 int16 = 647
        switch "a" {
        case "a":
            fmt.Println(v3)
        default:
            if ((!(v1 == int64(v9))) || true) {
                var v10 int8 = -63
                _ = v10
            }
        }
        _ = v9
    }
    var v11 int16 = 1
    _ = len(v5)
    v4 = v4
    if false {
        var v12 string = ("a" + "")
        var v13 uint = 706
        var v14 float64 = (v3 / v3)
        _ = v12
        _ = v13
        _ = v14
    } else {
        var v15 float32 = v4
        v16 := GenS1{F0: v6, F1: v11}
        _ = v15
        _ = v16
    }
    var v17 float64 = 3.14159
    v18 := GenS1{F0: 198, F1: v11}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
