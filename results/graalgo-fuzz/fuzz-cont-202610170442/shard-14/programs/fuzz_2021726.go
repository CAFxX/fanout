package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint64 }
type GenS2 struct { F0 int16; F1 int32; F2 int32 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int16; F1 int }
func genFn0(p0 string, p1 uint32) uint {
    for fi5 := 0; fi5 < 3; fi5++ {
        p1 = p1
        v6 := []int{fi5}
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p0)
    var v7 int16 = -32768
    _ = p0
    _ = p1
    _ = v7
    return 646
}
func genFn1(p0 uint16) uint8 {
    var v8 int64 = -129
    p0 = 65535
    _ = p0
    _ = v8
    return 178
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -1
    var v3 float32 = (0.0 * 1.5)
    var v4 int16 = -1
    v3 -= -2.25
    switch "a" {
    case "\n\t":
        if (65535 > 1) {
            fmt.Println(v4, v2)
            v9 := GenS1{F0: 907, F1: 0, F2: 0}
            v3 = 1.5
            _ = v9
        }
    case "a":
        v10 := GenS1{F0: 0, F1: -824, F2: 4294967295}
        _ = v10
    default:
        var v11 int32 = 1
        _ = v11
    }
    v12 := GenS2{F0: v4, F1: -304, F2: 1}
    if false {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 string = ((("" + "\n\t") + "a") + "")
            var v15 uint8 = 42
            v15 = genFn1(0)
            _ = fi13
            _ = v14
            _ = v15
        }
        v3 = -2.25
        var v16 int64 = 42
        _ = v16
    }
    fmt.Println(v3, v2)
    v17 := []int{-9223372036854775808, 127, 1, -9223372036854775808}
    if ((((v3 * float32(-2.25)) * -2.25) + v3) >= -2.25) {
        fmt.Println(v3, v3)
        v18 := GenS2{F0: 42, F1: -337, F2: -2147483648}
        v2 = v2
        _ = v18
    }
    v19 := GenS3{F0: -2147483648, F1: (2 <= 4294967295)}
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v20 := GenS3{F0: 990, F1: true}
    v21 := GenS4{F0: v4, F1: 9223372036854775807}
    var v22 uint8 = 127
    var v23 int64 = 470
    fmt.Println(v4, v23, v2)
    v4 = (v4 / ((1000) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
