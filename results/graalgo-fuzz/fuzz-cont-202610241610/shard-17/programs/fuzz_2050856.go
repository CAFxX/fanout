package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint, p1 string, p2 uint16) string {
    switch -9223372036854775808 {
    case 32767:
        var v5 int64 = 9223372036854775807
        _ = v5
    }
    fmt.Println(p0)
    p0 = (^(p0))
    v6 := GenS1{F0: -634}
    if (0 < 1) {
        v7 := GenS1{F0: 361}
        v8 := GenS1{F0: -2147483648}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int8 = 1
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ""
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v4 float64 = 3.14159
    if false {
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS2{F0: 42, F1: 42, F2: 0}
            _ = fi11
            _ = v12
        }
    } else {
        v4 = 3.14159
    }
    if false {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(fi13)
            _ = fi13
        }
        if false {
            v4 -= ((v4 + v4) * -2.25)
        }
    } else {
        var v14 uint8 = 114
        _ = v14
    }
    if false {
        v4 -= v4
        v4 = (float64((float32(float64((3.14159 * 1e10))) - 0.1)) + (0.1 / v4))
    } else {
        fmt.Println(v4, v4)
        v3[2147483648] = 2
    }
    v15 := []int{-898, 127}
    v1 = v1
    v1 = ("a" + "hello")
    for fi16 := 0; fi16 < 2; fi16++ {
        v4 -= v4
        _ = fi16
    }
    v17 := GenS3{F0: 2, F1: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
