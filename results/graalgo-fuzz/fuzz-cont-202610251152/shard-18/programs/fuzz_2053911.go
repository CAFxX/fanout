package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint8, p2 int) string {
    p1 += p0
    for fi4 := 0; fi4 < 12; fi4++ {
        if false {
            var v5 uint = 4294967295
            var v6 uint16 = 1
            _ = v5
            _ = v6
        } else {
            v7 := GenS1{F0: 0, F1: p1}
            p2 = 32767
            _ = v7
        }
        _ = fi4
    }
    p1 = (p1 * 42)
    _ = p0
    _ = p1
    _ = p2
    return "a"
}
func genFn1(p0 bool, p1 int8, p2 int64) uint8 {
    var v8 int = 9223372036854775807
    if (p0 || (0.1 <= 1.5)) {
        if ((((-2.25 + (-2.25 - float64(3.14159))) - float64(1e10)) * 0.1) > 1.5) {
            var v9 int16 = -347
            var v10 uint32 = 1
            v8 = 128
            _ = v10
            _ = v9
        }
        v11 := []int{-1, (v8 | v8), 119, v8}
        _ = v11
    } else {
        p0 = p0
    }
    var v12 float64 = (1e10 + 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v8
    return 127
}

func main() {
    v1 := []int{-324, -9223372036854775808, 2147483647, -2147483648}
    var v2 uint16 = 816
    var v3 float64 = (3.14159 * (((0.0 + 1e10) / (1e10 * 0.1)) / 0.1))
    var v13 uint64 = 323
    var v14 uint32 = 42
    var v15 uint = 18446744073709551615
    var v16 bool = true
    v17 := GenS2{F0: (v14 / (((v14 - 2)) | 1))}
    v3 = 0.0
    var v18 uint32 = 695
    v19 := GenS2{F0: 0}
    v15 -= (v15 & v15)
    _ = len(v1)
    v16 = (v16 && ((v2 << ((573) & 63)) <= v2))
    if v16 {
        if (!(false && ("foo bar" != "a"))) {
            var v20 int = 42
            var v21 float64 = float64(0.0)
            _ = v20
            _ = v21
        }
    } else {
        v22 := GenS1{F0: 418, F1: genFn1(v16, 98, -9223372036854775808)}
        fmt.Println(v18, v14)
        v13 = 18446744073709551615
        _ = v22
    }
    v13 += (v13 ^ (v13 * 2))
    v23 := GenS2{F0: (v18 << (((v15 / ((255) | 1))) & 63))}
    var v24 float64 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
