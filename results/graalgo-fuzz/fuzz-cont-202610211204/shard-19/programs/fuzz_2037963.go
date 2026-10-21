package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int32 }
func genFn0(p0 bool, p1 uint8, p2 uint32) string {
    if (!(-1 == 1)) {
        v5 := []int{9223372036854775807, 2147483647}
        p0 = (9223372036854775807 != -809)
        _ = v5
    }
    switch "foo bar" {
    case "a":
        var v6 int32 = -2147483648
        _ = v6
    case "":
        fmt.Println(p2, p2, p0)
    case "hello":
        v7 := GenS1{F0: 2147483647, F1: 1, F2: 42}
        _ = v7
    }
    var v8 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ""
}
func genFn1(p0 uint64, p1 int64) uint16 {
    v9 := []int{9223372036854775807}
    p0 = p0
    var v10 uint = 467
    if ((p0 >> (((v10 % ((42) | 1))) & 63)) > (p0 * p0)) {
        if len(v9) > 0 {
            v9[0] = -649
        }
        var v11 int = 4294967295
        _ = v11
    }
    v12 := GenS1{F0: 1, F1: 623, F2: -433}
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    _ = v9
    return 769
}

func main() {
    v1 := []uint32{2147483648, 65535, 2147483647}
    var v2 int16 = -443
    v3 := []int{-812, -1, 0, -1}
    var v4 uint16 = 255
    v2 = v2
    var v13 uint8 = 128
    v14 := GenS1{F0: 42, F1: 18446744073709551615, F2: -181}
    for fi15 := 0; fi15 < 10; fi15++ {
        v16 := GenS1{F0: 0, F1: 4294967295, F2: -271}
        _ = fi15
        _ = v16
    }
    var v17 uint8 = uint8(v4)
    var v18 uint8 = (v13 + v17)
    v2 = (v2 % ((1) | 1))
    v4 = v4
    if true {
        v1 = append(v1, 145)
        if true {
            v3 = append(v3, 762)
            v2 = 42
        }
    } else {
        fmt.Println(v2, v4, v4)
        v4 = (v4 / ((0) | 1))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
