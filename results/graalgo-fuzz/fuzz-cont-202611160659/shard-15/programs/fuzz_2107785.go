package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 uint32) uint {
    v6 := GenS1{F0: 504, F1: p0, F2: 4294967295}
    v7 := []int{-558, 127, 0}
    _ = p0
    _ = v6
    _ = v7
    return 503
}

func main() {
    var v1 bool = true
    var v2 int32 = 1
    v3 := []uint32{42, 4294967295}
    var v4 uint = 112
    var v5 uint8 = 136
    v5 = (v5 >> ((42) & 63))
    var v8 int32 = v2
    v9 := GenS2{F0: 65535}
    var v10 uint8 = (v5 + v5)
    if v1 {
        var v11 int16 = 0
        switch ("" + ((("\n\t" + ("\n\t" + "a")) + "foo bar") + (((("hello" + ("a" + "\n\t")) + ("foo bar" + "a")) + "") + ("" + ("\n\t" + "foo bar"))))) {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = 65535
            }
        case "a":
            var v12 bool = (-128 <= -128)
            _ = v12
        }
        v5 += v5
        _ = v11
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 int64 = 1
        var v15 int16 = 839
        _ = fi13
        _ = v14
        _ = v15
    }
    v10 += 204
    if (4294967295 < 4294967295) {
        var v16 int16 = 1000
        if (v16 >= v16) {
            v2 = (v8 << (((v4 - v4)) & 63))
        }
        _ = v16
    } else {
        v3 = append(v3, 6)
    }
    if ((!false) && false) {
        if (v2 < (v2 | (v8 * (v8 % ((v2) | 1))))) {
            var v17 float64 = -2.25
            v5 = uint8(v4)
            _ = v17
        } else {
            _ = len(v3)
        }
        v18 := GenS1{F0: 0, F1: 4294967295, F2: 9223372036854775807}
        fmt.Println(v10, v8, v1)
        _ = v18
    }
    v19 := GenS2{F0: 480}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if ((v4 | v4) < (v4 / ((0) | 1))) {
        if (42 == -9223372036854775808) {
            v20 := GenS3{F0: v4, F1: -32768}
            if (1 != 42) {
                v8 = v2
                fmt.Println(v10, v8, v4)
                v2 = -193
            }
            fmt.Println(v2, v10, v10)
            _ = v20
        }
        v21 := GenS2{F0: 2147483648}
        _ = v21
    }
    var v22 uint = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
