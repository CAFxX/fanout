package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
func genFn0(p0 uint8) int16 {
    p0 -= 2
    var v7 uint16 = 42
    var v8 uint64 = 0
    switch "" {
    case "\n\t":
        v9 := []int{971}
        _ = v9
    case "hello":
        var v10 uint64 = v8
        _ = v10
    default:
        var v11 uint8 = p0
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    return int16(v8)
}
func genFn1(p0 int8, p1 int16) string {
    v12 := GenS1{F0: 255, F1: 179}
    for fi13 := 0; fi13 < 6; fi13++ {
        if true {
            v14 := []int{fi13, fi13, fi13, -1}
            v15 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615}
            var v16 bool = false
            _ = v14
            _ = v15
            _ = v16
        } else {
            p0 = (p0 % (((p0 & -128)) | 1))
            p0 -= int8(fi13)
            var v17 int = (fi13 << ((1) & 63))
            _ = v17
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v12
    return ((("a" + "foo bar") + "\n\t") + ("\n\t" + ""))
}

func main() {
    var v1 int16 = 776
    var v2 int16 = (v1 & (v1 >> ((195) & 63)))
    var v3 uint32 = 4294967295
    var v4 int = -9223372036854775808
    var v5 int16 = v1
    var v6 int32 = -2147483648
    if (v3 >= (v3 + v3)) {
        if true {
            var v18 uint16 = 42
            _ = v18
        }
        v19 := GenS1{F0: 0, F1: 47}
        _ = v19
    }
    v1 = v2
    var v20 uint64 = 319
    var v21 uint8 = 37
    v4 = v4
    v22 := []int{}
    if (100 >= -16) {
        if len(v22) > 0 {
            v22[0] = 238
        }
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v1 = (v5 - (v1 * (v2 >> ((1) & 63))))
        _ = fi23
    }
    _ = len(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
