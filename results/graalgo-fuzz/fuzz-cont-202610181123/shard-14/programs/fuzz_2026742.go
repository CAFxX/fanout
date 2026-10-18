package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 int; F2 uint32 }

func main() {
    var v1 int8 = -128
    var v2 uint = 42
    v3 := make(map[string]uint32)
    v2 -= (v2 * v2)
    v2 += (^(v2))
    v4 := []int{9223372036854775807, 9223372036854775807}
    v4 = append(v4, 32768)
    v1 -= (v1 ^ v1)
    if (0 < 32767) {
        fmt.Println(v2, v1)
    }
    v1 += (v1 % ((1) | 1))
    delete(v3, "a")
    if ((v1 + v1) >= 0) {
        var v5 uint32 = 221
        fmt.Println(v5)
        _ = v5
    }
    if ((false || (2147483647 >= 681)) || true) {
        var v6 float32 = ((0.0 * 3.14159) * ((3.14159 + -2.25) + float32(0.1)))
        v1 = (v1 & -1)
        _ = v6
    } else {
        v7 := GenS1{F0: 128, F1: int(v2)}
        fmt.Println(v2, v1)
        _ = v7
    }
    if ((-2.25 != 3.14159) && (false && true)) {
        if ((!false) || (1 >= 127)) {
            v2 = 2
            var v8 int32 = -2147483648
            _ = v8
        }
        if ((1e10 + (1e10 + (1e10 + 1e10))) == (((0.1 - (1.5 + 0.1)) - 0.1) / 1e10)) {
            v9 := v3["a"]
            v10 := GenS2{F0: 2, F1: 330, F2: v9}
            _ = v10
            _ = v9
        } else {
            v11 := GenS2{F0: 42, F1: -1, F2: 4294967295}
            _ = v11
        }
    } else {
        var v12 int16 = 1000
        v13 := GenS1{F0: 1, F1: 815}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v12
        _ = v13
    }
    var v14 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
