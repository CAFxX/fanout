package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 bool }
func genFn0(p0 uint) int {
    if true {
        var v6 int32 = 916
        if (42 != 0) {
            var v7 int8 = 96
            _ = v7
        } else {
            v8 := []int{42, 2147483647}
            if len(v8) > 0 {
                v8[0] = -2147483648
            }
            switch ("a" + "") {
            case "foo bar":
                fmt.Println(v6)
            case "\n\t":
                v9 := GenS1{F0: "a", F1: 18446744073709551615, F2: (1 >= -2147483648)}
                _ = v9
            }
            _ = v8
        }
        _ = v6
    } else {
        p0 -= p0
    }
    p0 = p0
    var v10 uint = 2
    var v11 int = 688
    p0 = (p0 % ((v10) | 1))
    _ = p0
    _ = v10
    _ = v11
    return v11
}
func genFn1(p0 string, p1 int16, p2 int64) bool {
    p0 = "a"
    v12 := []int{-295, 1, -2147483648, genFn0(1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return true
}

func main() {
    var v1 uint16 = 255
    var v2 uint8 = 127
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    var v5 int = 128
    v13 := []int{genFn0(18446744073709551615), (-(v5))}
    var v14 int32 = 2147483647
    _ = len(v3)
    var v15 float64 = float64(3.14159)
    v4["foo bar"] = 0
    fmt.Println(v14, v14)
    var v16 uint16 = 2
    if len(v13) > 0 {
        v13[0] = (v5 * genFn0(888))
    }
    switch "" {
    case "hello":
        v14 -= (v14 >> ((uint(v1)) & 63))
    case "\n\t":
        var v17 bool = (false && (0.0 > 0.0))
        _ = v17
    case "":
        v18 := v3["foo bar"]
        _ = v18
    default:
        var v19 float64 = 1e10
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
