package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 string; F1 uint32; F2 int }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 int16, p1 string, p2 int32) uint32 {
    var v7 uint16 = 0
    var v8 uint64 = 4294967295
    v8 -= 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1() uint8 {
    var v9 uint16 = 42
    var v10 string = "a"
    v11 := []int{-215}
    var v12 int = 4294967295
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 93
}

func main() {
    var v1 uint64 = 2
    v2 := []int{-9223372036854775808}
    var v3 uint64 = (v1 ^ (v1 ^ 901))
    v4 := make(map[int]string)
    var v5 float64 = 0.1
    var v6 float64 = v5
    for fi13 := 0; fi13 < 12; fi13++ {
        _ = len(v4)
        v14 := GenS1{F0: v3, F1: (!(-71 < 42))}
        _ = fi13
        _ = v14
    }
    switch ("" + "\n\t") {
    case "\n\t":
        var v15 uint32 = 42
        _ = v15
    case "foo bar":
        if true {
            v5 += float64((1.5 + (3.14159 - -2.25)))
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 float64 = ((v6 * v5) + (3.14159 * v6))
                fmt.Println(v1, fi16)
                v18 := GenS1{F0: (v3 - v1), F1: (-2147483648 < 1)}
                _ = fi16
                _ = v17
                _ = v18
            }
            v6 += v6
        }
    default:
        var v19 int = 4294967295
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        if false {
            v21 := v4[fi20]
            _ = v21
        } else {
            fmt.Println(fi20)
            var v22 int8 = -83
            var v23 int64 = -9223372036854775808
            _ = v22
            _ = v23
        }
        var v24 uint64 = (v1 | 42)
        _ = fi20
        _ = v24
    }
    var v25 int64 = -9223372036854775808
    fmt.Println(v25, v1)
    fmt.Println(v5, v6)
    v26 := GenS2{F0: "hello", F1: 4294967295, F2: 714}
    var v27 uint16 = 130
    fmt.Println(v1, v5, v1)
    v5 = float64(float32(v6))
    v28 := GenS3{F0: 42, F1: 4294967295}
    v29 := GenS4{F0: 128, F1: (^(v3))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
