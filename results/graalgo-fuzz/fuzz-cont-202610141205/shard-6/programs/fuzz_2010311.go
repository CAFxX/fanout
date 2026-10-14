package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int64; F1 string }
func genFn0(p0 uint, p1 int32) uint32 {
    var v5 int64 = -9223372036854775808
    v5 += (v5 + (v5 * (v5 - v5)))
    v5 = 42
    var v6 string = "foo bar"
    p1 += (p1 / ((int32(p0)) | 1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483648
}
func genFn1() uint16 {
    v7 := []int{}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint16 = 841
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(fi8)
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    var v11 uint32 = 171
    var v12 uint64 = 4294967295
    _ = v11
    _ = v12
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 1
    var v2 int8 = 127
    var v3 float64 = 1e10
    v4 := make(map[string]bool)
    var v13 int64 = -9223372036854775808
    delete(v4, "foo bar")
    v14 := GenS1{F0: -720}
    v15 := GenS1{F0: 0}
    fmt.Println(v1, v13, v2)
    v16 := GenS2{F0: 4294967295}
    v13 -= v13
    v17 := GenS2{F0: 255}
    v18 := GenS3{F0: v1, F1: 2147483648}
    if (!(((v2 > (v2 << ((4294967295) & 63))) || false) || false)) {
        v19 := GenS4{F0: v13, F1: "\n\t"}
        fmt.Println(v1, v1, v1)
        _ = v19
    } else {
        var v20 uint32 = 2147483647
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS1{F0: -2147483648}
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
