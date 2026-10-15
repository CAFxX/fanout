package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int16; F2 string }
func genFn0(p0 int16, p1 int64, p2 bool) uint8 {
    var v7 int64 = p1
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint8 = 42
        var v10 float32 = (3.14159 - 3.14159)
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint32 = 2147483648
    var v12 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 int32, p2 int32) int {
    for fi13 := 0; fi13 < 10; fi13++ {
        p2 = (p2 / (((p0 % ((42) | 1))) | 1))
        _ = fi13
    }
    if (0 < 2) {
        fmt.Println(p2, p1)
        fmt.Println(p1, p1, p1)
    }
    var v14 int16 = -514
    for fi15 := 0; fi15 < 6; fi15++ {
        if true {
            v16 := GenS1{F0: v14}
            _ = v16
        }
        _ = fi15
    }
    v17 := []int{1}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v17
    return 127
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{42}
    var v3 uint64 = 569
    v4 := []string{"a", "", "hello", ""}
    var v5 bool = false
    v6 := make(map[int]bool)
    v5 = false
    if ((!v5) || (((0.1 - 1e10) <= 0.0) || (0 >= genFn0(-355, 9223372036854775807, (v3 >= v3))))) {
        v3 = v3
        v5 = (v5 && (false || v5))
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v1, v3)
        switch "" {
        case "\n\t":
            v1 = uint(v3)
        case "foo bar":
            fmt.Println(v5, v5, v5)
        default:
            v18 := GenS2{F0: 678}
            _ = v18
        }
    }
    v3 = 711
    var v19 bool = (0 != 9223372036854775807)
    fmt.Println(v1, v19, v19)
    v1 += 796
    if (v19 || (!(0 >= 65535))) {
        for fi20 := 0; fi20 < 8; fi20++ {
            if v5 {
                v21 := GenS1{F0: 1}
                _ = v21
            }
            _ = fi20
        }
        fmt.Println(v5, v3, v3)
        v22 := GenS3{F0: v5, F1: -32768, F2: "foo bar"}
        _ = v22
    }
    if v5 {
        var v23 uint64 = (v3 ^ 0)
        v1 = (v1 ^ 42)
        _ = v23
    } else {
        var v24 int8 = -128
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        v26 := GenS1{F0: -542}
        _ = len(v6)
        _ = fi25
        _ = v26
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
