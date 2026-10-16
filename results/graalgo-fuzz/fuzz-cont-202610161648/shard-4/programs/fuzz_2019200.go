package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 string; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint; F2 uint64 }
type GenS5 struct { F0 int64; F1 int64; F2 uint8 }
func genFn0(p0 uint16) int16 {
    var v5 int16 = 517
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 float64 = 3.14159
    switch 118 {
    case 32767:
        v8 := GenS1{F0: 81, F1: v5}
        _ = v8
    case 4294967295:
        v9 := GenS1{F0: 1, F1: v5}
        _ = v9
    case -400:
        v10 := GenS1{F0: 9, F1: -1}
        _ = v10
    default:
        v11 := []int{-830, 128}
        _ = v11
    }
    _ = p0
    _ = v5
    _ = v7
    return (-(v5))
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]bool)
    var v3 uint32 = 0
    var v4 int = 9223372036854775807
    var v12 int16 = genFn0(44)
    v13 := GenS2{F0: "hello", F1: 4294967295, F2: -128}
    fmt.Println(v12, v1, v4)
    v14 := []int{v4, (v4 * (v4 << ((138) & 63))), v4, 128}
    var v15 uint = 65535
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := v2[v4]
        fmt.Println(v12)
        var v18 uint32 = v3
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: (^(v3))}
    fmt.Println(v12, v15, v12)
    v4 = v4
    var v20 uint = 4294967295
    switch ((("hello" + "\n\t") + "") + ("foo bar" + (("\n\t" + "") + "\n\t"))) {
    case "foo bar":
        var v21 int64 = -9223372036854775808
        _ = v21
    case "a":
        v22 := GenS3{F0: v3}
        _ = v22
    case "\n\t":
        v15 = (v20 - (v20 ^ (v15 & 117)))
    default:
        v23 := GenS3{F0: v3}
        _ = v23
    }
    var v24 uint = v20
    var v25 int32 = 217
    v26 := GenS4{F0: 9223372036854775807, F1: 0, F2: 4294967295}
    v24 = 198
    var v27 int64 = 42
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 uint8 = 0
        v15 -= v20
        if (v12 < (v12 >> ((65535) & 63))) {
            v30 := GenS5{F0: 1, F1: 90, F2: v29}
            v31 := GenS3{F0: 0}
            _ = v30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
