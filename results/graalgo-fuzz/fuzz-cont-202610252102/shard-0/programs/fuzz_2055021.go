package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }
func genFn0() int32 {
    switch "foo bar" {
    case "\n\t":
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 uint8 = 127
            _ = fi5
            _ = v6
        }
    default:
        var v7 bool = ((!true) && true)
        _ = v7
    }
    if (1 == 42) {
        var v8 int64 = -784
        var v9 int8 = -1
        _ = v8
        _ = v9
    } else {
        var v10 uint = 0
        fmt.Println(v10)
        var v11 uint32 = 1
        _ = v10
        _ = v11
    }
    var v12 float32 = (float32(1.5) * 0.0)
    var v13 int32 = 1
    _ = v12
    _ = v13
    return v13
}
func genFn1(p0 uint16) int16 {
    var v14 float32 = 1e10
    v14 -= 1e10
    switch "a" {
    case "a":
        switch "\n\t" {
        case "a":
            var v15 uint16 = p0
            _ = v15
        default:
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 int8 = -1
                v17 -= v17
                v17 += (v17 & (^(v17)))
                _ = fi16
                _ = v17
            }
        }
    case "hello":
        switch 128 {
        case -9223372036854775808:
            fmt.Println(v14, v14)
        case 2147483647:
            fmt.Println(v14, p0)
        }
    default:
        v18 := GenS1{F0: "", F1: 42}
        _ = v18
    }
    _ = p0
    _ = v14
    return 0
}

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 65535
    var v3 uint8 = 128
    var v4 float32 = 1.5
    var v19 uint16 = v2
    v20 := GenS2{F0: v3}
    v2 = (v2 - 913)
    if true {
        v21 := GenS3{F0: 9223372036854775807}
        _ = v21
    } else {
        var v22 float64 = 0.0
        _ = v22
    }
    fmt.Println(v2)
    if (-9223372036854775808 <= 0) {
        var v23 uint16 = v2
        var v24 float32 = (0.1 - v4)
        _ = v23
        _ = v24
    } else {
        v1 += (v1 - v1)
        v25 := []int{32767, -2147483648, 817}
        v26 := GenS4{F0: 1, F1: (v2 * v2), F2: uint64(v19)}
        _ = v25
        _ = v26
    }
    switch 4294967295 {
    case -2147483648:
        switch -9223372036854775808 {
        case 2147483647:
            var v27 int64 = -9223372036854775808
            _ = v27
        case -222:
            v1 = v1
        case 303:
            v28 := GenS1{F0: "hello", F1: -38}
            _ = v28
        default:
            v29 := []int{-9223372036854775808, -1, 176}
            _ = v29
        }
    case 832:
        v30 := GenS2{F0: 128}
        _ = v30
    }
    v31 := []int{128, 436, 291, 616}
    switch 0 {
    case 0:
        v3 = (^(v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
}
